package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ak-ansari/go-pdf-service/model"
)

type StudentAPIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewStudentAPIClient(baseURL string) *StudentAPIClient {
	return &StudentAPIClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *StudentAPIClient) GetStudentByID(id string) (*model.Student, error) {
	url := fmt.Sprintf("%s/students/%s", c.BaseURL, id)

	// will add header as per suggested changes for authentication across service to service call
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch student, status: %d", resp.StatusCode)
	}

	var student model.Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, err
	}

	return &student, nil
}
