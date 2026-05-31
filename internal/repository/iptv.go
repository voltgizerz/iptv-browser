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
	GetCategories(context.Context) ([]model.Category, error)
	GetChannels(context.Context) ([]model.Channel, error)
	GetStreams(context.Context) ([]model.Stream, error)
	GetLogos(context.Context) ([]model.Logo, error)
	GetStreamURL(context.Context, string) (string, error)
}

type iptvRepository struct {
	client *http.Client

	mu sync.RWMutex

	countries  []model.Country
	categories []model.Category
	channels   []model.Channel
	streams    []model.Stream
	logos      []model.Logo

	streamMap map[string]string
	logoMap   map[string]string

	lastRefresh time.Time
}

const cacheTTL = 6 * time.Hour

func NewIPTVRepository() IPTVRepository {
	return &iptvRepository{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		streamMap: make(map[string]string),
		logoMap:   make(map[string]string),
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

func (r *iptvRepository) ensureCache(_ context.Context) error {

	r.mu.RLock()

	if len(r.channels) > 0 &&
		time.Since(r.lastRefresh) < cacheTTL {
		r.mu.RUnlock()
		return nil
	}

	r.mu.RUnlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.channels) > 0 &&
		time.Since(r.lastRefresh) < cacheTTL {
		return nil
	}

	var (
		countries  []model.Country
		categories []model.Category
		channels   []model.Channel
		streams    []model.Stream
		logos      []model.Logo

		countriesErr  error
		categoriesErr error
		channelsErr   error
		streamsErr    error
		logosErr      error
	)

	var wg sync.WaitGroup

	wg.Add(5)

	go func() {
		defer wg.Done()

		countriesErr = r.fetch(
			"https://iptv-org.github.io/api/countries.json",
			&countries,
		)
	}()

	go func() {
		defer wg.Done()

		categoriesErr = r.fetch(
			"https://iptv-org.github.io/api/categories.json",
			&categories,
		)
	}()

	go func() {
		defer wg.Done()

		channelsErr = r.fetch(
			"https://iptv-org.github.io/api/channels.json",
			&channels,
		)
	}()

	go func() {
		defer wg.Done()

		streamsErr = r.fetch(
			"https://iptv-org.github.io/api/streams.json",
			&streams,
		)
	}()

	go func() {
		defer wg.Done()

		logosErr = r.fetch(
			"https://iptv-org.github.io/api/logos.json",
			&logos,
		)
	}()

	wg.Wait()

	if countriesErr != nil {
		return countriesErr
	}

	if categoriesErr != nil {
		return categoriesErr
	}

	if channelsErr != nil {
		return channelsErr
	}

	if streamsErr != nil {
		return streamsErr
	}

	if logosErr != nil {
		return logosErr
	}

	streamMap := make(map[string]string)

	for _, stream := range streams {
		streamMap[stream.Channel] = stream.URL
	}

	logoMap := make(map[string]string)

	for _, logo := range logos {
		logoMap[logo.Channel] = logo.URL
	}

	r.countries = countries
	r.categories = categories
	r.channels = channels
	r.streams = streams
	r.logos = logos

	r.streamMap = streamMap
	r.logoMap = logoMap

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

func (r *iptvRepository) GetCategories(ctx context.Context) ([]model.Category, error) {
	if err := r.ensureCache(ctx); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.categories, nil
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

func (r *iptvRepository) GetStreamURL(ctx context.Context, channelID string) (string, error) {
	if err := r.ensureCache(ctx); err != nil {
		return "", err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.streamMap[channelID], nil
}
