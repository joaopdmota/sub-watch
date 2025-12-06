# Database Design

This document serves as the source of truth for the database schema.
We use Mermaid to visualize the relationships before implementing them in Go/SQL.

## Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    USERS ||--o{ SUBSCRIPTIONS : "owns"
    USERS ||--o{ CATEGORIES : "manages"
    USERS ||--o{ PAYMENT_METHODS : "manages"
    
    CATEGORIES ||--o{ SUBSCRIPTIONS : "classifies"
    PAYMENT_METHODS ||--o{ SUBSCRIPTIONS : "pays_for"
    SUBSCRIPTIONS ||--o{ PRICE_HISTORY : "has"

    USERS {
        uuid id PK
        string name
        string email
        string password_hash
        timestamp created_at
        timestamp updated_at
    }

    CATEGORIES {
        uuid id PK
        uuid user_id FK "Owner (nullable for system defaults)"
        string name
        string icon "Icon identifier"
        string color "Hex code"
    }

    PAYMENT_METHODS {
        uuid id PK
        uuid user_id FK
        string name "e.g. Nubank Card"
        string type "CREDIT_CARD, DEBIT, BOLETO, PIX"
        timestamp created_at
    }

    SUBSCRIPTIONS {
        uuid id PK
        uuid user_id FK
        uuid category_id FK
        uuid payment_method_id FK
        string service_name
        float price
        string currency "BRL, USD"
        string cycle "MONTHLY, YEARLY, WEEKLY"
        date next_billing_date
        string status "ACTIVE, PAUSED, CANCELED"
        text notes "General notes"
        timestamp created_at
        timestamp updated_at
    }

    PRICE_HISTORY {
        uuid id PK
        uuid subscription_id FK
        float old_price
        float new_price
        timestamp changed_at
        text reason "e.g. Annual adjustment"
    }
```

## Implementation Plan
1. Define entities in `application/domain`.
2. Create migration scripts in `infra/database/migrations`.
3. Implement Repositories in `infra/database/repositories`.
