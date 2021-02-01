package honeycomb

import (
	"log"
	"os"

	"github.com/honeycombio/beeline-go"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
)

var _ monitoring.Monitoring = &Monitoring{}

type Monitoring struct {
}

// Init sets up monitoring
func Init(serviceName string) {
	key := os.Getenv("HONEYCOMB_API_KEY")
	if len(key) == 0 {
		log.Println("No API key was defined for Honeycomb, telemetry will not be sent. Set the HONEYCOMB_API_KEY and HONEYCOMB_DATASET env vars to enable.")
		return
	}
	dataset := os.Getenv("HONEYCOMB_DATASET")
	if len(dataset) == 0 {
		log.Println("No dataset name was defined for Honeycomb, telemetry will not be sent. Set the HONEYCOMB_API_KEY and HONEYCOMB_DATASET env vars to enable.")
		return
	}
	beeline.Init(beeline.Config{
		WriteKey:    key,
		Dataset:     dataset,
		ServiceName: serviceName,
	})
	log.Printf("Initialized Honeycomb for dataset %q", dataset)

	monitoring.Default = &Monitoring{}
}

func (m *Monitoring) Close() error {
	beeline.Close()
	return nil
}
