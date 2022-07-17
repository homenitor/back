package entities

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/values"
)

type ProbeListingView struct {
	ID   string
	Name string
}

type GetSamplesView struct {
	MeasuredAt time.Time
	Values     map[string]float64
	Average    float64
}

type Probe struct {
	id           string
	name         string
	humidities   []*Sample
	temperatures []*Sample
}

func NewProbeWithID(id string) *Probe {
	return &Probe{
		id:           id,
		humidities:   make([]*Sample, 0),
		temperatures: make([]*Sample, 0),
	}
}

func NewProbeWithIdAndName(id string, name string) *Probe {
	return &Probe{
		id:           id,
		name:         name,
		humidities:   make([]*Sample, 0),
		temperatures: make([]*Sample, 0),
	}
}

func (p *Probe) ID() string {
	return p.id
}

func (p *Probe) Name() string {
	return p.name
}

func (p *Probe) RecordHumidity(s *Sample) {
	p.humidities = append(p.humidities, s)
}

func (p *Probe) SamplesByCategory(category values.SampleCategory) []*Sample {
	switch category {
	case values.HUMIDITY_SAMPLE_CATEGORY:
		return p.humidities
	case values.TEMPERATURE_SAMPLE_CATEGORY:
		return p.temperatures
	}

	return nil
}

func (p *Probe) LatestHumidity() (*Sample, error) {
	humiditiesLength := len(p.humidities)
	hasAtLeastOneHumidity := humiditiesLength > 0
	if !hasAtLeastOneHumidity {
		return nil, common.ErrNoSampleValueInProbe
	}

	lastHumidityIndex := humiditiesLength - 1

	return p.humidities[lastHumidityIndex], nil
}

func (p *Probe) RecordTemperature(s *Sample) {
	p.temperatures = append(p.temperatures, s)
}

func (p *Probe) LatestTemperature() (*Sample, error) {
	temperaturesLength := len(p.temperatures)
	hasAtLeastOneTemperature := temperaturesLength > 0
	if !hasAtLeastOneTemperature {
		return nil, common.ErrNoSampleValueInProbe
	}

	lastTemperaturesIndex := temperaturesLength - 1

	return p.temperatures[lastTemperaturesIndex], nil
}
