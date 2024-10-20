package pools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sieugene/jp-reader/handlers"
	"github.com/sieugene/jp-reader/internal/database"
)

type Project struct {
	Images  []string  `json:"images"`
	Name    string    `json:"name"`
	OcrData []OcrData `json:"ocrData"`
}

type OcrData struct {
	Data OcrDataContent `json:"data"`
	Name string         `json:"name"`
}

type OcrDataContent struct {
	Blocks    []Block `json:"blocks"`
	ImgHeight int     `json:"img_height"`
	ImgWidth  int     `json:"img_width"`
	Version   string  `json:"version"`
}

type Block struct {
	Box         []float64     `json:"box"`
	FontSize    int           `json:"font_size"`
	Lines       []string      `json:"lines"`
	LinesCoords [][][]float64 `json:"lines_coords"`
	Vertical    bool          `json:"vertical"`
}

type Projects struct {
	Projects []Project `json:"projects"`
}

func StartPollingProjects(apiCfq handlers.ApiConfig) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := pollProjects(apiCfq); err != nil {
				log.Printf("Error polling projects: %v", err)
			}
		}
	}
}

func pollProjects(apiCfq handlers.ApiConfig) error {
	MOKURO_SERVICE_ENDPOINT := os.Getenv("MOKURO_SERVICE")
	requestEndpoint := MOKURO_SERVICE_ENDPOINT + "/projects"

	resp, err := http.Get(requestEndpoint)
	if err != nil {
		return fmt.Errorf("failed to fetch projects: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var result = Projects{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to parse projects JSON: %w", err)
	}

	existingProjects, err := apiCfq.DB.GetProjects(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get existing projects: %w", err)
	}

	existingMap := make(map[string]bool)
	for _, project := range existingProjects {
		existingMap[project.Name] = true
	}

	for _, project := range result.Projects {
		if existingMap[project.Name] {
			delete(existingMap, project.Name)
			continue
		}

		rawData, err := json.Marshal(project.OcrData)
		if err != nil {
			log.Printf("Failed to marshal OCR data for project %s: %v", project.Name, err)
			return err
		}

		ocrData := json.RawMessage(rawData)

		if _, err := apiCfq.DB.CreateProject(context.Background(), database.CreateProjectParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdateAt:  time.Now().UTC(),
			Name:      project.Name,
			Images:    project.Images,
			OcrData:   ocrData,
		}); err != nil {
			log.Printf("Failed to create project %s: %v", project.Name, err)
		}
	}

	for name := range existingMap {
		if err := apiCfq.DB.DeleteProjectByName(context.Background(), name); err != nil {
			log.Printf("Failed to delete project %s: %v", name, err)
		}
	}

	return nil
}
