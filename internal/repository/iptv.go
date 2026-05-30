package repository

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/voltgizerz/iptv-browser/internal/model"
)

type IPTVRepository interface {
	GetCountries(context.Context) ([]model.Country, error)
	GetChannels(context.Context) ([]model.Channel, error)
	GetStreams(context.Context) ([]model.Stream, error)
	GetLogos(context.Context) ([]model.Logo, error)
}

type iptvRepository struct {
	client *http.Client

	mu sync.RWMutex

	countries []model.Country
	channels  []model.Channel
	streams   []model.Stream
	logos     []model.Logo

	lastRefresh time.Time
}

const cacheTTL = 1 * time.Hour

func NewIPTVRepository() IPTVRepository {
	return &iptvRepository{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (r *iptvRepository) fetch(url string, target any) error {
	resp, err := r.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func (r *iptvRepository) ensureCache(ctx context.Context) error {

	r.mu.RLock()

	if time.Since(r.lastRefresh) < cacheTTL &&
		len(r.channels) > 0 {
		r.mu.RUnlock()
		return nil
	}

	r.mu.RUnlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	if time.Since(r.lastRefresh) < cacheTTL &&
		len(r.channels) > 0 {
		return nil
	}

	var countries []model.Country
	var channels []model.Channel
	var streams []model.Stream
	var logos []model.Logo

	if err := r.fetch(
		"https://iptv-org.github.io/api/countries.json",
		&countries,
	); err != nil {
		return err
	}

	if err := r.fetch(
		"https://iptv-org.github.io/api/channels.json",
		&channels,
	); err != nil {
		return err
	}

	if err := r.fetch(
		"https://iptv-org.github.io/api/streams.json",
		&streams,
	); err != nil {
		return err
	}

	if err := r.fetch(
		"https://iptv-org.github.io/api/logos.json",
		&logos,
	); err != nil {
		return err
	}

	r.countries = countries
	r.channels = channels
	r.streams = streams
	r.logos = logos
	r.lastRefresh = time.Now()

	return nil
}

func (r *iptvRepository) GetCountries(ctx context.Context) ([]model.Country, error) {

	if err := r.ensureCache(ctx); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.countries, nil
}

func (r *iptvRepository) GetChannels(ctx context.Context) ([]model.Channel, error) {

	if err := r.ensureCache(ctx); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.channels, nil
}

func (r *iptvRepository) GetStreams(ctx context.Context) ([]model.Stream, error) {

	if err := r.ensureCache(ctx); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.streams, nil
}

func (r *iptvRepository) GetLogos(ctx context.Context) ([]model.Logo, error) {

	if err := r.ensureCache(ctx); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.logos, nil
}