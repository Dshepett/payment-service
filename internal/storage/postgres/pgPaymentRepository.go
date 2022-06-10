package postgres

import (
	"database/sql"

	"github.com/Dshepett/payment-service/internal/models"
	_ "github.com/lib/pq"
)

type PgPaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PgPaymentRepository {
	return &PgPaymentRepository{
		db: db,
	}
}

func (p *PgPaymentRepository) AddPayment(payment *models.Payment) error {
	query := "INSERT INTO payments(user_id, user_email, amount, currency, created_at, updated_at, status)" +
		" VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;"
	err := p.db.QueryRow(query, payment.UserId, payment.UserEmail,
		payment.Amount, payment.Currency, payment.CreatedAt, payment.UpdatedAt, payment.Status).Scan(&payment.Id)
	return err
}

func (p *PgPaymentRepository) UpdatePaymentStatus(payment *models.Payment) error {
	query := "UPDATE payments SET status = $1 WHERE id = $2;"
	_, err := p.db.Exec(query, payment.Status, payment.Id)
	return err
}

func (p *PgPaymentRepository) GetPayment(id int) (*models.Payment, error) {
	query := "SELECT * FROM payments WHERE id = $1"
	payment := models.Payment{}
	err := p.db.QueryRow(query, id).Scan(&payment.Id, &payment.UserId, &payment.UserEmail, &payment.Amount,
		&payment.Currency, &payment.CreatedAt, &payment.UpdatedAt, &payment.Status)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (p *PgPaymentRepository) GetPaymentsByUserId(id int) ([]models.Payment, error) {
	query := "SELECT * FROM payments WHERE user_id = $1"
	var payments []models.Payment
	rows, err := p.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		payment := models.Payment{}
		err := rows.Scan(&payment.Id, &payment.UserId, &payment.UserEmail, &payment.Amount,
			&payment.Currency, &payment.CreatedAt, &payment.UpdatedAt, &payment.Status)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (p *PgPaymentRepository) GetPaymentsByUserEmail(email string) ([]models.Payment, error) {
	query := "SELECT * FROM payments WHERE user_email = $1"
	var payments []models.Payment
	rows, err := p.db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		payment := models.Payment{}
		err := rows.Scan(&payment.Id, &payment.UserId, &payment.UserEmail, &payment.Amount,
			&payment.Currency, &payment.CreatedAt, &payment.UpdatedAt, &payment.Status)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (p *PgPaymentRepository) DeletePaymentById(id int) error {
	query := "DELETE FROM payments WHERE id=$1"
	_, err := p.db.Exec(query, id)
	return err
}

func (p *PgPaymentRepository) Close() error {
	return p.db.Close()
}
