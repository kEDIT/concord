package crawler

import (
	"context"
	"errors"
	"fmt"
)

type result struct {
	url   string
	depth int
}

type Crawler struct {
	pool    int          //Default = 1
	resc    chan *result //Default capacity 1 result struct
	errc    chan error
	ctx     context.Context //Carry from main, use for cancelation propegation
	maxdep  int
	filters []string
}

//Constructor & Configuration

func NewCrawler() *Crawler {
	return &Crawler{
		pool:    1,
		resc:    make(chan *result, 1),
		errc:    make(chan error, 1),
		ctx:     context.Background(),
		maxdep:  3,
		filters: []string{},
	}
}

func (c *Crawler) Set(opt ...func(crawler *Crawler) error) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Set Option ", r)
		}
	}()

	for _, o := range opt {
		if err := o(c); err != nil {
			panic(err)
		}
	}

	return nil
}

func Poolsize(n int) func(c *Crawler) error {
	return func(c *Crawler) error {
		if err := c.setPoolSize(n); err != nil {
			return err
		}
		return nil
	}
}

func (c *Crawler) setPoolSize(n int) error {
	if n < 1 {
		return errors.New("Invalid worker pool size specified.")
	}
	c.pool = n
	return nil
}

func ResultChanCap(n int) func(c *Crawler) error {
	return func(c *Crawler) error {
		if err := c.setResCap(n); err != nil {
			return err
		}
		return nil
	}
}

func (c *Crawler) setResCap(n int) error {
	//TODO: handling error should account for problems assigning chan to crawler
	ch := make(chan *result, n)
	c.resc = ch
	return nil
}

func Context(ctx context.Context) func(c *Crawler) error {
	return func(c *Crawler) error {
		if err := c.setContext(ctx); err != nil {
			return err
		}
		return nil
	}
}

func (c *Crawler) setContext(ctx context.Context) error {
	//TODO:  handle error
	c.ctx = ctx
	return nil
}

func MaxDepth(n int) func(c *Crawler) error {
	return func(c *Crawler) error {
		return c.setMaxDepth(n)
	}
}

func (c *Crawler) setMaxDepth(n int) error {
	if n <= 0 {
		return errors.New("Invalid parameter. Set: MaxDepth must be integer >= 1")
	}
	c.maxdep = n
	return nil
}

func Filters(filters []string) func(c *Crawler) error {
	return func(c *Crawler) error {
		return c.setFilters(filters)
	}
}

func (c *Crawler) setFilters(filters []string) error {
	if len(filters) == 0 {
		return errors.New("Insufficient number of arguments passed to setFilter method.")
	}
	c.filters = filters
	return nil
}
