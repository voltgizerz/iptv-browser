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

func (s *IPTVService) GetChannels(
	ctx context.Context,
	country string,
	search string,
) ([]model.ChannelResponse, error) {

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

	result := make([]model.ChannelResponse, 0)

	for _, ch := range channels {

		if country != "" &&
			!strings.EqualFold(ch.Country, country) {
			continue
		}

		if search != "" &&
			!strings.Contains(
				strings.ToLower(ch.Name),
				strings.ToLower(search),
			) {
			continue
		}

		if !streamMap[ch.ID] {
			continue
		}

		result = append(result, model.ChannelResponse{
			ID:      ch.ID,
			Name:    ch.Name,
			Country: ch.Country,
			Logo:    logoMap[ch.ID],
		})
	}

	return result, nil
}

func (s *IPTVService) GetStream(
	ctx context.Context,
	id string,
) (string, error) {

	streams, err := s.repo.GetStreams(ctx)
	if err != nil {
		return "", err
	}

	for _, stream := range streams {
		if stream.Channel == id {
			return stream.URL, nil
		}
	}

	return "", nil
}