package job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/category"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/payment"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/thumbnail"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/transaction"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/user"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/utils"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type MessageOpenAI struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIChatOpenAIRequestBody struct {
	Model            string          `json:"model"`
	Temperature      int             `json:"temperature"`
	MaxTokens        int32           `json:"max_tokens"`
	TopP             int             `json:"top_p"`
	FrequencyPenalty int             `json:"frequency_penalty"`
	PresencePenalty  int             `json:"presence_penalty"`
	Messages         []MessageOpenAI `json:"messages"`
}

type ChoiceOpenAI struct {
	Message MessageOpenAI `json:"message"`
}

type APIChatOpenAIResponseBody struct {
	Choices []ChoiceOpenAI `json:"choices"`
}

type CustomerService struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Job struct {
	ID             uuid.UUID `gorm:"type:varchar(100);"`
	Title          string    `gorm:"type:varchar(255);not null"`
	Description    string    `gorm:"type:text;not null"`
	RewardEarned   float64   `gorm:"type:decimal;not null"`
	FromAddressID  uuid.UUID `gorm:"type:varchar(100)"`
	FromAddress    address.Address
	ToAddressID    uuid.UUID `gorm:"type:varchar(100)"`
	ToAddress      address.Address
	Status         string    `gorm:"type:varchar(50);not null"`
	HelperRequired uint      `gorm:"not null"`
	CategoryID     uuid.UUID `gorm:"type:varchar(100);not null"`
	Category       category.Category
	UserID         uuid.UUID `gorm:"type:varchar(100);not null"`
	User           user.User
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Transactions   []transaction.Transaction
	Thumbnails     []thumbnail.Thumbnail
	DeletedAt      gorm.DeletedAt
}

func FromUseCase(job *entities.Job) *Job {
	jobTransactions := make([]transaction.Transaction, len(job.Transactions))
	for i, _transaction := range job.Transactions {
		jobTransactions[i] = transaction.Transaction{
			ID:       _transaction.ID,
			UserID:   _transaction.UserID,
			JobID:    _transaction.JobID,
			Type:     _transaction.Type,
			SubTotal: _transaction.SubTotal,
			Tax:      _transaction.Tax,
			Total:    _transaction.Total,
			Payment: payment.Payment{
				ID:            _transaction.Payment.ID,
				ExternalID:    _transaction.Payment.ExternalID,
				Status:        _transaction.Payment.Status,
				TransactionID: _transaction.Payment.TransactionID,
				Amount:        _transaction.Payment.Amount,
				InvoiceURL:    _transaction.Payment.InvoiceURL,
			},
			CreatedAt: _transaction.CreatedAt,
			UpdatedAt: _transaction.UpdatedAt,
		}
	}
	jobThumbnails := make([]thumbnail.Thumbnail, len(job.Thumbnails))
	for i, _thumbnail := range job.Thumbnails {
		jobThumbnails[i] = thumbnail.Thumbnail{
			ID:          _thumbnail.ID,
			ImageKey:    _thumbnail.ImageKey,
			JobID:       job.ID,
			Description: _thumbnail.Description,
		}
	}

	return &Job{
		ID:             job.ID,
		Title:          job.Title,
		Description:    job.Description,
		RewardEarned:   job.RewardEarned,
		UserID:         job.UserID,
		FromAddressID:  job.FromAddress.ID,
		FromAddress:    address.Address{ID: job.FromAddress.ID},
		ToAddressID:    job.ToAddress.ID,
		ToAddress:      address.Address{ID: job.ToAddress.ID},
		Status:         job.Status,
		HelperRequired: job.HelperRequired,
		Category:       category.Category{ID: job.Category.ID},
		CreatedAt:      job.CreatedAt,
		UpdatedAt:      job.UpdatedAt,
		Transactions:   jobTransactions,
		Thumbnails:     jobThumbnails,
	}
}

func (j *Job) ToUseCase() *entities.Job {
	jobTransactions := make([]entities.Transaction, len(j.Transactions))
	for i, _transaction := range j.Transactions {
		jobTransactions[i] = entities.Transaction{
			ID:       _transaction.ID,
			Type:     _transaction.Type,
			UserID:   _transaction.UserID,
			JobID:    _transaction.JobID,
			SubTotal: _transaction.SubTotal,
			Tax:      _transaction.Tax,
			Total:    _transaction.Total,
			Payment: entities.Payment{
				ID:            _transaction.Payment.ID,
				ExternalID:    _transaction.Payment.ExternalID,
				Status:        _transaction.Payment.Status,
				TransactionID: _transaction.Payment.TransactionID,
				Amount:        _transaction.Payment.Amount,
				InvoiceURL:    _transaction.Payment.InvoiceURL,
			},
			CreatedAt: _transaction.CreatedAt,
			UpdatedAt: _transaction.UpdatedAt,
		}
	}
	jobThumbnails := make([]entities.Thumbnail, len(j.Thumbnails))
	for i, _thumbnail := range j.Thumbnails {
		jobThumbnails[i] = entities.Thumbnail{
			ID:          _thumbnail.ID,
			ImageKey:    _thumbnail.ImageKey,
			Description: _thumbnail.Description,
		}
	}

	return &entities.Job{
		ID:           j.ID,
		Title:        j.Title,
		UserID:       j.UserID,
		Description:  j.Description,
		RewardEarned: j.RewardEarned,
		FromAddress: entities.Address{
			ID:        j.FromAddress.ID,
			Address:   j.FromAddress.Address,
			City:      j.FromAddress.City,
			State:     j.FromAddress.State,
			Country:   j.FromAddress.Country,
			ZipCode:   j.FromAddress.ZipCode,
			Longitude: j.FromAddress.Longitude,
			Latitude:  j.FromAddress.Latitude,
			CreatedAt: j.FromAddress.CreatedAt,
			UpdatedAt: j.FromAddress.UpdatedAt,
		},
		ToAddress: entities.Address{
			ID:        j.ToAddress.ID,
			Address:   j.ToAddress.Address,
			City:      j.ToAddress.City,
			State:     j.ToAddress.State,
			Country:   j.ToAddress.Country,
			ZipCode:   j.ToAddress.ZipCode,
			Longitude: j.ToAddress.Longitude,
			Latitude:  j.ToAddress.Latitude,
			CreatedAt: j.ToAddress.CreatedAt,
			UpdatedAt: j.ToAddress.UpdatedAt,
		},
		Status:         j.Status,
		HelperRequired: j.HelperRequired,
		Category: entities.Category{
			ID:   j.Category.ID,
			Name: j.Category.Name,
		},
		CreatedAt:    j.CreatedAt,
		UpdatedAt:    j.UpdatedAt,
		Transactions: jobTransactions,
		Thumbnails:   jobThumbnails,
	}
}

func CustomerServiceFromUseCase(cs *entities.JobCustomerService) *CustomerService {
	return &CustomerService{
		Question: cs.Question,
		Answer:   cs.Answer,
	}
}

func (cs *CustomerService) Talk() (*entities.JobCustomerService, error) {
	var client = &http.Client{}
	var requestBody = APIChatOpenAIRequestBody{
		Model:            "gpt-3.5-turbo",
		Temperature:      1,
		MaxTokens:        255,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Messages: []MessageOpenAI{
			{Role: "system", Content: "You're a product manager who knows all the product, and you're customer service for that product. The product briefs is platform connects users with helpers for moving, delivering, shopping, or recycling tasks. A platform that connects individuals who need help with daily tasks (customers) with freelancers (helpers). Platform name is Find Helpers."},
			{Role: "user", Content: "The question is" + cs.Question},
		},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(requestBody)
	if err != nil {
		fmt.Println("json.NewEncoder(&buf).Encode(requestBody)", err)
		return &entities.JobCustomerService{}, err
	}
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", &buf)
	if err != nil {
		fmt.Println("http.NewRequest", err)
		return &entities.JobCustomerService{}, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_API_KEY"))

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do", err)
		return &entities.JobCustomerService{}, err
	}
	defer response.Body.Close()

	responseJSON := APIChatOpenAIResponseBody{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("json.NewDecoder(response.Body).Decode(&responseJSON)", err)
		return &entities.JobCustomerService{}, err
	}

	fmt.Println("responseJSON", utils.PrettyPrint(responseJSON))

	return &entities.JobCustomerService{
		Question: cs.Question,
		Answer:   responseJSON.Choices[0].Message.Content,
	}, err
}
