package rabbit

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"sync/atomic"
	"time"
)

type Client struct {
	// Connection
	conn *amqp.Connection
	ch   *amqp.Channel

	// Reconnect
	connEstablished *atomic.Bool

	// Config
	cfg    Config
	logger *zerolog.Logger
}

func Connect(ctx context.Context, cfg Config) (*Client, error) {
	rabbitCfg := amqp.Config{
		Vhost:     cfg.Vhost,
		Heartbeat: cfg.Heartbeat,
	}

	conn, err := amqp.DialConfig(cfg.Host, rabbitCfg)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.Qos(cfg.Prefetch, 0, false)
	if err != nil {
		return nil, err
	}

	connEstablished := new(atomic.Bool)
	connEstablished.Store(false)

	client := &Client{
		conn:            conn,
		ch:              ch,
		connEstablished: connEstablished,
		logger:          zerolog.Ctx(ctx),
		cfg:             cfg,
	}

	go client.checkReconnect(ctx)

	return client, nil
}

func (c *Client) SetupQueue(exchange, queue, routing, kind string, durable, exclusive, autoDelete, internal, noWait bool, args amqp.Table) error {
	if err := c.ExchangeDeclare(exchange, kind, durable, autoDelete, internal, noWait, args); err != nil {
		return err
	}

	if _, err := c.QueueDeclare(queue, durable, autoDelete, exclusive, noWait, args); err != nil {
		return err
	}

	if err := c.QueueBind(queue, routing, exchange, noWait, args); err != nil {
		return err
	}

	return nil
}

func (c *Client) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return c.ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (c *Client) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return c.ch.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
}

func (c *Client) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	return c.ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

func (c *Client) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return c.ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (c *Client) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	return c.ch.QueueBind(name, key, exchange, noWait, args)
}

func (c *Client) Ack(tag uint64, multiple bool) error {
	return c.ch.Ack(tag, multiple)
}

func (c *Client) Nack(tag uint64, multiple, requeue bool) error {
	return c.ch.Nack(tag, multiple, requeue)
}

func (c *Client) checkReconnect(ctx context.Context) {
	errChan := make(chan *amqp.Error)
	c.conn.NotifyClose(errChan)

	for err := range errChan {
		c.connEstablished.Store(false)
		c.Close()

		c.logger.Error().Err(err).Msg("received error from rabbit, trying to reconnect...")

		for {
			if newClient, err := Connect(ctx, c.cfg); err == nil {
				c.connEstablished.Store(true)

				c.conn = newClient.conn
				c.ch = newClient.ch

				c.logger.Info().Msg("successfully reconnected to rabbit")

				break
			} else {
				c.logger.Error().Err(err).Msg("error reconnection to rabbit, retrying...")
			}

			time.Sleep(c.cfg.ReconnectDuration)
		}
	}
}

func (c *Client) Close() error {
	errs := make([]error, 0, 2)

	if err := c.closeConn(); err != nil {
		errs = append(errs, err)
	}

	if err := c.closeChan(); err != nil {
		errs = append(errs, err)
	}

	err := errors.Join(errs...)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) closeConn() error {
	if c.conn == nil {
		return nil
	}
	if err := c.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (c *Client) closeChan() error {
	if c.ch == nil {
		return nil
	}
	if err := c.ch.Close(); err != nil {
		return err
	}
	return nil
}

func (c *Client) isConnEstablished() bool {
	return c.connEstablished.Load()
}
