package service

import (
	"context"
	"strings"

	"github.com/voltgizerz/iptv-browser/internal/model"
	"github.com/voltgizerz/iptv-browser/internal/repository"
)

type IPTVService struct {
	repo repository.IPTVRepository
}

func NewIPTVService(
	repo repository.IPTVRepository,
) *IPTVService {
	return &IPTVService{
		repo: repo,
	}
}

func (s *IPTVService) GetCountries(
	ctx context.Context,
) ([]model.Country, error) {
	return s.repo.GetCountries(ctx)
}

func (s *IPTVService) GetCategories(
	ctx context.Context,
) ([]model.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *IPTVService) GetChannels(
	ctx context.Context,
	country string,
	category string,
	search string,
) ([]model.ChannelResponse, error) {
	category = strings.ToLower(strings.TrimSpace(category))
	search = strings.ToLower(strings.TrimSpace(search))

	channels, err := s.repo.GetChannels(ctx)
	if err != nil {
		return nil, err
	}

	streams, err := s.repo.GetStreams(ctx)
	if err != nil {
		return nil, err
	}

	logos, err := s.repo.GetLogos(ctx)
	if err != nil {
		return nil, err
	}

	streamMap := map[string]bool{}
	for _, s := range streams {
		streamMap[s.Channel] = true
	}

	logoMap := map[string]string{}
	for _, l := range logos {
		logoMap[l.Channel] = l.URL
	}

	result := make([]model.ChannelResponse, 0, len(channels))

	for _, ch := range channels {

		if country != "" &&
			!strings.EqualFold(ch.Country, country) {
			continue
		}

		if category != "" &&
			!matchesCategory(ch.Categories, category) {
			continue
		}

		if search != "" && !matchesChannelSearch(ch, search) {
			continue
		}

		if !streamMap[ch.ID] {
			continue
		}

		result = append(result, model.ChannelResponse{
			ID:         ch.ID,
			Name:       ch.Name,
			Country:    ch.Country,
			Categories: ch.Categories,
			Logo:       logoMap[ch.ID],
			IsNSFW:     ch.IsNSFW,
		})
	}

	return result, nil
}

func matchesCategory(categories []string, category string) bool {
	for _, c := range categories {
		if strings.Contains(strings.ToLower(c), category) {
			return true
		}
	}

	return false
}

func matchesChannelSearch(ch model.Channel, search string) bool {
	values := []string{
		ch.ID,
		ch.Name,
		ch.Country,
		ch.Network,
	}

	values = append(values, ch.AltNames...)
	values = append(values, ch.Categories...)

	for _, value := range values {
		if strings.Contains(strings.ToLower(value), search) {
			return true
		}
	}

	return false
}

func (s *IPTVService) GetStream(
	ctx context.Context,
	id string,
) (string, error) {
	return s.repo.GetStreamURL(ctx, id)
}
