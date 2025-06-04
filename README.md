**SONSAWAN NGAMSOM**

email: sonsawan.n.th@gmail.com
tel: 062-608-1634

### Playtorium - Take Home Assignment**

A checkout system implementation in Go that handles product purchases with complex coupon and discount management.

### Coupon Management
- Support for multiple coupon types:
    - Regular coupons
    - On-top discounts
    - Seasonal promotions
- Flexible calculation modes:
    - Fixed amount discounts
    - Percentage-based discounts
    - Category-specific discounts
    - Point-based discounts
    - Buy X Get Y discounts

### Cart Management
- Shopping cart functionality
- Product categorization
- Total price calculation
- Discount application
- Multiple items support

## System Requirements

- Go 1.23.7 or higher

## Installation
- clone this repo [GITHUB](https://github.com/murphy6867/playtorium_take_home_assignment.git) 
- download dependency: `go mod download`
- Setup .env with .env.example
- Setup database and connection string
- run migrating: `go run migration/migrate.go`

## Project Structure
internal/
    ├── app/
    │ ├── cart/
    │ ├── cart_item/
    │ ├── coupon/
    │ └── applied_coupon/
    └── utils/

## Usage

### Applying Coupons

The system supports different types of coupons with various calculation modes:

1. **Regular Coupons**
    - Fixed amount discounts
    - Percentage-based discounts

2. **On-top Discounts**
    - Category-specific percentage discounts
    - Point-based discounts

3. **Seasonal Discounts**
    - Buy X Get Y discount type

### Discount Calculation

Discounts are calculated based on:
- Coupon type priority (regular → on-top → seasonal)
- Calculation mode
- Cart total
- Product categories (for category-specific discounts)
- Points (for point-based discounts)

## Error Handling

The system implements robust error handling with domain-specific errors:
- Invalid coupon types
- Invalid calculation modes
- Missing discount data
- Point usage validation
- Cart and coupon validation

## API Documentation
