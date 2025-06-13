# Payment System Test Report - Fri 13 Jun 2025 03:02:15 PM PKT

## Domain Layer Test Results
=== RUN   TestEmail
=== RUN   TestEmail/create_valid_email
=== RUN   TestEmail/create_email_with_uppercase
=== RUN   TestEmail/create_email_with_whitespace
=== RUN   TestEmail/cannot_create_email_with_empty_string
=== RUN   TestEmail/cannot_create_email_with_invalid_format
=== RUN   TestEmail/get_domain
=== RUN   TestEmail/get_username
=== RUN   TestEmail/email_equality
=== RUN   TestEmail/email_string_representation
--- PASS: TestEmail (0.00s)
    --- PASS: TestEmail/create_valid_email (0.00s)
    --- PASS: TestEmail/create_email_with_uppercase (0.00s)
    --- PASS: TestEmail/create_email_with_whitespace (0.00s)
    --- PASS: TestEmail/cannot_create_email_with_empty_string (0.00s)
    --- PASS: TestEmail/cannot_create_email_with_invalid_format (0.00s)
    --- PASS: TestEmail/get_domain (0.00s)
    --- PASS: TestEmail/get_username (0.00s)
    --- PASS: TestEmail/email_equality (0.00s)
    --- PASS: TestEmail/email_string_representation (0.00s)
=== RUN   TestShippingAddress
=== RUN   TestShippingAddress/create_valid_shipping_address
=== RUN   TestShippingAddress/cannot_create_address_with_empty_customer_ID
=== RUN   TestShippingAddress/cannot_create_address_with_empty_first_name
=== RUN   TestShippingAddress/cannot_create_address_with_empty_last_name
=== RUN   TestShippingAddress/cannot_create_address_with_empty_address_line_1
=== RUN   TestShippingAddress/cannot_create_address_with_empty_city
=== RUN   TestShippingAddress/cannot_create_address_with_empty_state
=== RUN   TestShippingAddress/cannot_create_address_with_empty_postal_code
=== RUN   TestShippingAddress/cannot_create_address_with_empty_country_code
=== RUN   TestShippingAddress/cannot_create_address_with_invalid_country_code
=== RUN   TestShippingAddress/cannot_create_address_with_invalid_phone
=== RUN   TestShippingAddress/cannot_create_address_with_invalid_label
=== RUN   TestShippingAddress/get_full_name
=== RUN   TestShippingAddress/get_full_address
=== RUN   TestShippingAddress/update_label
=== RUN   TestShippingAddress/cannot_update_to_invalid_label
=== RUN   TestShippingAddress/set_and_unset_default
=== RUN   TestShippingAddress/validate_for_shipping
--- PASS: TestShippingAddress (0.00s)
    --- PASS: TestShippingAddress/create_valid_shipping_address (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_customer_ID (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_first_name (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_last_name (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_address_line_1 (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_city (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_state (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_postal_code (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_empty_country_code (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_invalid_country_code (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_invalid_phone (0.00s)
    --- PASS: TestShippingAddress/cannot_create_address_with_invalid_label (0.00s)
    --- PASS: TestShippingAddress/get_full_name (0.00s)
    --- PASS: TestShippingAddress/get_full_address (0.00s)
    --- PASS: TestShippingAddress/update_label (0.00s)
    --- PASS: TestShippingAddress/cannot_update_to_invalid_label (0.00s)
    --- PASS: TestShippingAddress/set_and_unset_default (0.00s)
    --- PASS: TestShippingAddress/validate_for_shipping (0.00s)
=== RUN   TestCustomerStatus
=== RUN   TestCustomerStatus/valid_statuses
=== RUN   TestCustomerStatus/invalid_status
--- PASS: TestCustomerStatus (0.00s)
    --- PASS: TestCustomerStatus/valid_statuses (0.00s)
    --- PASS: TestCustomerStatus/invalid_status (0.00s)
=== RUN   TestNewCustomer
=== RUN   TestNewCustomer/create_valid_customer
=== RUN   TestNewCustomer/cannot_create_customer_with_invalid_email
=== RUN   TestNewCustomer/cannot_create_customer_with_empty_first_name
=== RUN   TestNewCustomer/cannot_create_customer_with_empty_last_name
=== RUN   TestNewCustomer/cannot_create_customer_with_invalid_phone
=== RUN   TestNewCustomer/create_customer_without_phone
--- PASS: TestNewCustomer (0.00s)
    --- PASS: TestNewCustomer/create_valid_customer (0.00s)
    --- PASS: TestNewCustomer/cannot_create_customer_with_invalid_email (0.00s)
    --- PASS: TestNewCustomer/cannot_create_customer_with_empty_first_name (0.00s)
    --- PASS: TestNewCustomer/cannot_create_customer_with_empty_last_name (0.00s)
    --- PASS: TestNewCustomer/cannot_create_customer_with_invalid_phone (0.00s)
    --- PASS: TestNewCustomer/create_customer_without_phone (0.00s)
=== RUN   TestCustomerEmailUpdate
=== RUN   TestCustomerEmailUpdate/update_email
=== RUN   TestCustomerEmailUpdate/cannot_update_to_invalid_email
--- PASS: TestCustomerEmailUpdate (0.00s)
    --- PASS: TestCustomerEmailUpdate/update_email (0.00s)
    --- PASS: TestCustomerEmailUpdate/cannot_update_to_invalid_email (0.00s)
=== RUN   TestCustomerPersonalInfoUpdate
=== RUN   TestCustomerPersonalInfoUpdate/update_personal_info
=== RUN   TestCustomerPersonalInfoUpdate/cannot_update_with_empty_first_name
=== RUN   TestCustomerPersonalInfoUpdate/cannot_update_with_empty_last_name
=== RUN   TestCustomerPersonalInfoUpdate/cannot_update_with_invalid_phone
--- PASS: TestCustomerPersonalInfoUpdate (0.00s)
    --- PASS: TestCustomerPersonalInfoUpdate/update_personal_info (0.00s)
    --- PASS: TestCustomerPersonalInfoUpdate/cannot_update_with_empty_first_name (0.00s)
    --- PASS: TestCustomerPersonalInfoUpdate/cannot_update_with_empty_last_name (0.00s)
    --- PASS: TestCustomerPersonalInfoUpdate/cannot_update_with_invalid_phone (0.00s)
=== RUN   TestCustomerStatusTransitions
=== RUN   TestCustomerStatusTransitions/activate_customer
=== RUN   TestCustomerStatusTransitions/cannot_activate_already_active_customer
=== RUN   TestCustomerStatusTransitions/deactivate_customer
=== RUN   TestCustomerStatusTransitions/cannot_deactivate_already_inactive_customer
=== RUN   TestCustomerStatusTransitions/suspend_customer
--- PASS: TestCustomerStatusTransitions (0.00s)
    --- PASS: TestCustomerStatusTransitions/activate_customer (0.00s)
    --- PASS: TestCustomerStatusTransitions/cannot_activate_already_active_customer (0.00s)
    --- PASS: TestCustomerStatusTransitions/deactivate_customer (0.00s)
    --- PASS: TestCustomerStatusTransitions/cannot_deactivate_already_inactive_customer (0.00s)
    --- PASS: TestCustomerStatusTransitions/suspend_customer (0.00s)
=== RUN   TestCustomerShippingAddresses
=== RUN   TestCustomerShippingAddresses/add_first_shipping_address
=== RUN   TestCustomerShippingAddresses/add_second_shipping_address
=== RUN   TestCustomerShippingAddresses/update_shipping_address
=== RUN   TestCustomerShippingAddresses/cannot_update_non-existent_address
=== RUN   TestCustomerShippingAddresses/remove_shipping_address
=== RUN   TestCustomerShippingAddresses/cannot_remove_only_shipping_address
=== RUN   TestCustomerShippingAddresses/removing_default_address_makes_first_remaining_address_default
=== RUN   TestCustomerShippingAddresses/set_default_shipping_address
=== RUN   TestCustomerShippingAddresses/cannot_set_non-existent_address_as_default
=== RUN   TestCustomerShippingAddresses/get_default_shipping_address
=== RUN   TestCustomerShippingAddresses/get_shipping_address_by_ID
=== RUN   TestCustomerShippingAddresses/cannot_get_non-existent_address
--- PASS: TestCustomerShippingAddresses (0.00s)
    --- PASS: TestCustomerShippingAddresses/add_first_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/add_second_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/update_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/cannot_update_non-existent_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/remove_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/cannot_remove_only_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/removing_default_address_makes_first_remaining_address_default (0.00s)
    --- PASS: TestCustomerShippingAddresses/set_default_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/cannot_set_non-existent_address_as_default (0.00s)
    --- PASS: TestCustomerShippingAddresses/get_default_shipping_address (0.00s)
    --- PASS: TestCustomerShippingAddresses/get_shipping_address_by_ID (0.00s)
    --- PASS: TestCustomerShippingAddresses/cannot_get_non-existent_address (0.00s)
=== RUN   TestCustomerQueryMethods
=== RUN   TestCustomerQueryMethods/query_methods_on_active_customer
=== RUN   TestCustomerQueryMethods/query_methods_on_customer_with_addresses
=== RUN   TestCustomerQueryMethods/inactive_customer_cannot_place_order
--- PASS: TestCustomerQueryMethods (0.00s)
    --- PASS: TestCustomerQueryMethods/query_methods_on_active_customer (0.00s)
    --- PASS: TestCustomerQueryMethods/query_methods_on_customer_with_addresses (0.00s)
    --- PASS: TestCustomerQueryMethods/inactive_customer_cannot_place_order (0.00s)
PASS
=== RUN   TestNewOrder
=== RUN   TestNewOrder/success_with_valid_inputs
=== RUN   TestNewOrder/error_with_empty_customer_ID
=== RUN   TestNewOrder/error_with_no_items
--- PASS: TestNewOrder (0.00s)
    --- PASS: TestNewOrder/success_with_valid_inputs (0.00s)
    --- PASS: TestNewOrder/error_with_empty_customer_ID (0.00s)
    --- PASS: TestNewOrder/error_with_no_items (0.00s)
=== RUN   TestOrderStatusTransitions
=== RUN   TestOrderStatusTransitions/mark_as_paid
=== RUN   TestOrderStatusTransitions/cannot_mark_as_paid_twice
=== RUN   TestOrderStatusTransitions/mark_as_fulfilled
=== RUN   TestOrderStatusTransitions/cannot_fulfill_unpaid_order
--- PASS: TestOrderStatusTransitions (0.00s)
    --- PASS: TestOrderStatusTransitions/mark_as_paid (0.00s)
    --- PASS: TestOrderStatusTransitions/cannot_mark_as_paid_twice (0.00s)
    --- PASS: TestOrderStatusTransitions/mark_as_fulfilled (0.00s)
    --- PASS: TestOrderStatusTransitions/cannot_fulfill_unpaid_order (0.00s)
=== RUN   TestOrderCancel
=== RUN   TestOrderCancel/cancel_created_order
=== RUN   TestOrderCancel/cancel_paid_order
=== RUN   TestOrderCancel/cannot_cancel_fulfilled_order
--- PASS: TestOrderCancel (0.00s)
    --- PASS: TestOrderCancel/cancel_created_order (0.00s)
    --- PASS: TestOrderCancel/cancel_paid_order (0.00s)
    --- PASS: TestOrderCancel/cannot_cancel_fulfilled_order (0.00s)
=== RUN   TestOrderItemManagement
=== RUN   TestOrderItemManagement/add_item_to_order
=== RUN   TestOrderItemManagement/cannot_add_item_after_payment
=== RUN   TestOrderItemManagement/remove_item_from_order
=== RUN   TestOrderItemManagement/cannot_remove_non-existent_item
=== RUN   TestOrderItemManagement/removing_last_item_cancels_order
=== RUN   TestOrderItemManagement/cannot_modify_order_after_fulfillment
--- PASS: TestOrderItemManagement (0.00s)
    --- PASS: TestOrderItemManagement/add_item_to_order (0.00s)
    --- PASS: TestOrderItemManagement/cannot_add_item_after_payment (0.00s)
    --- PASS: TestOrderItemManagement/remove_item_from_order (0.00s)
    --- PASS: TestOrderItemManagement/cannot_remove_non-existent_item (0.00s)
    --- PASS: TestOrderItemManagement/removing_last_item_cancels_order (0.00s)
    --- PASS: TestOrderItemManagement/cannot_modify_order_after_fulfillment (0.00s)
=== RUN   TestCalculateTotalAmount
=== RUN   TestCalculateTotalAmount/calculate_total_with_multiple_items
=== RUN   TestCalculateTotalAmount/order_must_have_at_least_one_item
=== RUN   TestCalculateTotalAmount/error_with_inconsistent_currencies
--- PASS: TestCalculateTotalAmount (0.00s)
    --- PASS: TestCalculateTotalAmount/calculate_total_with_multiple_items (0.00s)
    --- PASS: TestCalculateTotalAmount/order_must_have_at_least_one_item (0.00s)
    --- PASS: TestCalculateTotalAmount/error_with_inconsistent_currencies (0.00s)
PASS
=== RUN   TestPaymentStatus
=== RUN   TestPaymentStatus/valid_statuses
=== RUN   TestPaymentStatus/invalid_status
=== RUN   TestPaymentStatus/completed_status
=== RUN   TestPaymentStatus/final_status
=== RUN   TestPaymentStatus/can_be_refunded
=== RUN   TestPaymentStatus/can_be_cancelled
--- PASS: TestPaymentStatus (0.00s)
    --- PASS: TestPaymentStatus/valid_statuses (0.00s)
    --- PASS: TestPaymentStatus/invalid_status (0.00s)
    --- PASS: TestPaymentStatus/completed_status (0.00s)
    --- PASS: TestPaymentStatus/final_status (0.00s)
    --- PASS: TestPaymentStatus/can_be_refunded (0.00s)
    --- PASS: TestPaymentStatus/can_be_cancelled (0.00s)
=== RUN   TestCryptoCurrency
=== RUN   TestCryptoCurrency/get_supported_cryptocurrencies
=== RUN   TestCryptoCurrency/get_cryptocurrency_by_symbol
=== RUN   TestCryptoCurrency/get_cryptocurrency_by_symbol_case_insensitive
=== RUN   TestCryptoCurrency/unsupported_cryptocurrency
=== RUN   TestCryptoCurrency/is_supported
=== RUN   TestCryptoCurrency/validate_amount
=== RUN   TestCryptoCurrency/cryptocurrency_properties
--- PASS: TestCryptoCurrency (0.00s)
    --- PASS: TestCryptoCurrency/get_supported_cryptocurrencies (0.00s)
    --- PASS: TestCryptoCurrency/get_cryptocurrency_by_symbol (0.00s)
    --- PASS: TestCryptoCurrency/get_cryptocurrency_by_symbol_case_insensitive (0.00s)
    --- PASS: TestCryptoCurrency/unsupported_cryptocurrency (0.00s)
    --- PASS: TestCryptoCurrency/is_supported (0.00s)
    --- PASS: TestCryptoCurrency/validate_amount (0.00s)
    --- PASS: TestCryptoCurrency/cryptocurrency_properties (0.00s)
=== RUN   TestPaymentMethod
=== RUN   TestPaymentMethod/create_valid_payment_method
=== RUN   TestPaymentMethod/cannot_create_payment_method_with_empty_wallet_address
=== RUN   TestPaymentMethod/cannot_create_payment_method_with_unsupported_crypto
=== RUN   TestPaymentMethod/payment_method_expiration
=== RUN   TestPaymentMethod/payment_method_time_until_expiry
=== RUN   TestPaymentMethod/validate_amount
=== RUN   TestPaymentMethod/payment_method_properties
--- PASS: TestPaymentMethod (0.00s)
    --- PASS: TestPaymentMethod/create_valid_payment_method (0.00s)
    --- PASS: TestPaymentMethod/cannot_create_payment_method_with_empty_wallet_address (0.00s)
    --- PASS: TestPaymentMethod/cannot_create_payment_method_with_unsupported_crypto (0.00s)
    --- PASS: TestPaymentMethod/payment_method_expiration (0.00s)
    --- PASS: TestPaymentMethod/payment_method_time_until_expiry (0.00s)
    --- PASS: TestPaymentMethod/validate_amount (0.00s)
    --- PASS: TestPaymentMethod/payment_method_properties (0.00s)
=== RUN   TestNewPayment
=== RUN   TestNewPayment/create_valid_payment
=== RUN   TestNewPayment/cannot_create_payment_with_empty_order_ID
=== RUN   TestNewPayment/cannot_create_payment_with_invalid_amount
=== RUN   TestNewPayment/cannot_create_payment_with_empty_currency
=== RUN   TestNewPayment/cannot_create_payment_with_unsupported_crypto
--- PASS: TestNewPayment (0.00s)
    --- PASS: TestNewPayment/create_valid_payment (0.00s)
    --- PASS: TestNewPayment/cannot_create_payment_with_empty_order_ID (0.00s)
    --- PASS: TestNewPayment/cannot_create_payment_with_invalid_amount (0.00s)
    --- PASS: TestNewPayment/cannot_create_payment_with_empty_currency (0.00s)
    --- PASS: TestNewPayment/cannot_create_payment_with_unsupported_crypto (0.00s)
=== RUN   TestPaymentCryptoAmount
=== RUN   TestPaymentCryptoAmount/update_crypto_amount
=== RUN   TestPaymentCryptoAmount/cannot_update_crypto_amount_on_final_payment
=== RUN   TestPaymentCryptoAmount/cannot_update_with_invalid_crypto_amount
--- PASS: TestPaymentCryptoAmount (0.00s)
    --- PASS: TestPaymentCryptoAmount/update_crypto_amount (0.00s)
    --- PASS: TestPaymentCryptoAmount/cannot_update_crypto_amount_on_final_payment (0.00s)
    --- PASS: TestPaymentCryptoAmount/cannot_update_with_invalid_crypto_amount (0.00s)
=== RUN   TestPaymentStatusTransitions
=== RUN   TestPaymentStatusTransitions/mark_as_confirming
=== RUN   TestPaymentStatusTransitions/cannot_mark_as_confirming_from_wrong_status
=== RUN   TestPaymentStatusTransitions/cannot_mark_as_confirming_with_empty_transaction_hash
=== RUN   TestPaymentStatusTransitions/update_confirmations
=== RUN   TestPaymentStatusTransitions/auto_confirm_with_enough_confirmations
=== RUN   TestPaymentStatusTransitions/manual_confirm_payment
=== RUN   TestPaymentStatusTransitions/cannot_confirm_already_confirmed_payment
=== RUN   TestPaymentStatusTransitions/mark_as_failed
=== RUN   TestPaymentStatusTransitions/mark_as_expired
--- PASS: TestPaymentStatusTransitions (0.00s)
    --- PASS: TestPaymentStatusTransitions/mark_as_confirming (0.00s)
    --- PASS: TestPaymentStatusTransitions/cannot_mark_as_confirming_from_wrong_status (0.00s)
    --- PASS: TestPaymentStatusTransitions/cannot_mark_as_confirming_with_empty_transaction_hash (0.00s)
    --- PASS: TestPaymentStatusTransitions/update_confirmations (0.00s)
    --- PASS: TestPaymentStatusTransitions/auto_confirm_with_enough_confirmations (0.00s)
    --- PASS: TestPaymentStatusTransitions/manual_confirm_payment (0.00s)
    --- PASS: TestPaymentStatusTransitions/cannot_confirm_already_confirmed_payment (0.00s)
    --- PASS: TestPaymentStatusTransitions/mark_as_failed (0.00s)
    --- PASS: TestPaymentStatusTransitions/mark_as_expired (0.00s)
=== RUN   TestPaymentCancellation
=== RUN   TestPaymentCancellation/cancel_pending_payment
=== RUN   TestPaymentCancellation/cancel_confirming_payment
=== RUN   TestPaymentCancellation/cannot_cancel_confirmed_payment
--- PASS: TestPaymentCancellation (0.00s)
    --- PASS: TestPaymentCancellation/cancel_pending_payment (0.00s)
    --- PASS: TestPaymentCancellation/cancel_confirming_payment (0.00s)
    --- PASS: TestPaymentCancellation/cannot_cancel_confirmed_payment (0.00s)
=== RUN   TestPaymentRefunds
=== RUN   TestPaymentRefunds/full_refund_confirmed_payment
=== RUN   TestPaymentRefunds/partial_refund_confirmed_payment
=== RUN   TestPaymentRefunds/cannot_refund_non-confirmed_payment
=== RUN   TestPaymentRefunds/cannot_refund_more_than_payment_amount
=== RUN   TestPaymentRefunds/set_refund_transaction_hash
--- PASS: TestPaymentRefunds (0.00s)
    --- PASS: TestPaymentRefunds/full_refund_confirmed_payment (0.00s)
    --- PASS: TestPaymentRefunds/partial_refund_confirmed_payment (0.00s)
    --- PASS: TestPaymentRefunds/cannot_refund_non-confirmed_payment (0.00s)
    --- PASS: TestPaymentRefunds/cannot_refund_more_than_payment_amount (0.00s)
    --- PASS: TestPaymentRefunds/set_refund_transaction_hash (0.00s)
=== RUN   TestPaymentValidation
=== RUN   TestPaymentValidation/validate_exact_amount
=== RUN   TestPaymentValidation/validate_amount_within_tolerance
=== RUN   TestPaymentValidation/insufficient_amount
=== RUN   TestPaymentValidation/excessive_amount
--- PASS: TestPaymentValidation (0.00s)
    --- PASS: TestPaymentValidation/validate_exact_amount (0.00s)
    --- PASS: TestPaymentValidation/validate_amount_within_tolerance (0.00s)
    --- PASS: TestPaymentValidation/insufficient_amount (0.00s)
    --- PASS: TestPaymentValidation/excessive_amount (0.00s)
=== RUN   TestPaymentExternalService
=== RUN   TestPaymentExternalService/set_now_payments_ID
=== RUN   TestPaymentExternalService/cannot_set_empty_now_payments_ID
=== RUN   TestPaymentExternalService/set_callback_URL
--- PASS: TestPaymentExternalService (0.00s)
    --- PASS: TestPaymentExternalService/set_now_payments_ID (0.00s)
    --- PASS: TestPaymentExternalService/cannot_set_empty_now_payments_ID (0.00s)
    --- PASS: TestPaymentExternalService/set_callback_URL (0.00s)
=== RUN   TestPaymentExpiration
=== RUN   TestPaymentExpiration/payment_expiration
=== RUN   TestPaymentExpiration/cannot_mark_expired_payment_as_confirming
--- PASS: TestPaymentExpiration (0.00s)
    --- PASS: TestPaymentExpiration/payment_expiration (0.00s)
    --- PASS: TestPaymentExpiration/cannot_mark_expired_payment_as_confirming (0.00s)
=== RUN   TestPaymentQueryMethods
=== RUN   TestPaymentQueryMethods/query_methods_on_pending_payment
=== RUN   TestPaymentQueryMethods/query_methods_on_confirmed_payment
--- PASS: TestPaymentQueryMethods (0.00s)
    --- PASS: TestPaymentQueryMethods/query_methods_on_pending_payment (0.00s)
    --- PASS: TestPaymentQueryMethods/query_methods_on_confirmed_payment (0.00s)
=== RUN   TestRequiredConfirmations
=== RUN   TestRequiredConfirmations/BTC_required_confirmations
=== RUN   TestRequiredConfirmations/ETH_required_confirmations
=== RUN   TestRequiredConfirmations/LTC_required_confirmations
=== RUN   TestRequiredConfirmations/BCH_required_confirmations
=== RUN   TestRequiredConfirmations/XRP_required_confirmations
=== RUN   TestRequiredConfirmations/DOGE_required_confirmations
--- PASS: TestRequiredConfirmations (0.00s)
    --- PASS: TestRequiredConfirmations/BTC_required_confirmations (0.00s)
    --- PASS: TestRequiredConfirmations/ETH_required_confirmations (0.00s)
    --- PASS: TestRequiredConfirmations/LTC_required_confirmations (0.00s)
    --- PASS: TestRequiredConfirmations/BCH_required_confirmations (0.00s)
    --- PASS: TestRequiredConfirmations/XRP_required_confirmations (0.00s)
    --- PASS: TestRequiredConfirmations/DOGE_required_confirmations (0.00s)
PASS
=== RUN   TestProductStatus
=== RUN   TestProductStatus/valid_statuses
=== RUN   TestProductStatus/invalid_status
=== RUN   TestProductStatus/can_be_ordered
--- PASS: TestProductStatus (0.00s)
    --- PASS: TestProductStatus/valid_statuses (0.00s)
    --- PASS: TestProductStatus/invalid_status (0.00s)
    --- PASS: TestProductStatus/can_be_ordered (0.00s)
=== RUN   TestInventory
=== RUN   TestInventory/create_valid_inventory
=== RUN   TestInventory/cannot_create_inventory_with_negative_quantity
=== RUN   TestInventory/cannot_create_inventory_with_negative_reserved
=== RUN   TestInventory/cannot_create_inventory_with_reserved_exceeding_total
=== RUN   TestInventory/reserve_stock
=== RUN   TestInventory/cannot_reserve_more_than_available
=== RUN   TestInventory/release_stock
=== RUN   TestInventory/fulfill_stock
=== RUN   TestInventory/check_low_stock
=== RUN   TestInventory/check_out_of_stock
--- PASS: TestInventory (0.00s)
    --- PASS: TestInventory/create_valid_inventory (0.00s)
    --- PASS: TestInventory/cannot_create_inventory_with_negative_quantity (0.00s)
    --- PASS: TestInventory/cannot_create_inventory_with_negative_reserved (0.00s)
    --- PASS: TestInventory/cannot_create_inventory_with_reserved_exceeding_total (0.00s)
    --- PASS: TestInventory/reserve_stock (0.00s)
    --- PASS: TestInventory/cannot_reserve_more_than_available (0.00s)
    --- PASS: TestInventory/release_stock (0.00s)
    --- PASS: TestInventory/fulfill_stock (0.00s)
    --- PASS: TestInventory/check_low_stock (0.00s)
    --- PASS: TestInventory/check_out_of_stock (0.00s)
=== RUN   TestCategory
=== RUN   TestCategory/create_valid_category
=== RUN   TestCategory/cannot_create_category_with_empty_name
=== RUN   TestCategory/create_category_with_parent
=== RUN   TestCategory/update_category_name
=== RUN   TestCategory/cannot_update_to_empty_name
=== RUN   TestCategory/set_and_remove_parent
=== RUN   TestCategory/cannot_set_self_as_parent
--- PASS: TestCategory (0.00s)
    --- PASS: TestCategory/create_valid_category (0.00s)
    --- PASS: TestCategory/cannot_create_category_with_empty_name (0.00s)
    --- PASS: TestCategory/create_category_with_parent (0.00s)
    --- PASS: TestCategory/update_category_name (0.00s)
    --- PASS: TestCategory/cannot_update_to_empty_name (0.00s)
    --- PASS: TestCategory/set_and_remove_parent (0.00s)
    --- PASS: TestCategory/cannot_set_self_as_parent (0.00s)
=== RUN   TestCategoryPath
=== RUN   TestCategoryPath/create_breadcrumb
=== RUN   TestCategoryPath/empty_path
--- PASS: TestCategoryPath (0.00s)
    --- PASS: TestCategoryPath/create_breadcrumb (0.00s)
    --- PASS: TestCategoryPath/empty_path (0.00s)
=== RUN   TestProduct
=== RUN   TestProduct/create_valid_product
=== RUN   TestProduct/cannot_create_product_with_empty_name
=== RUN   TestProduct/cannot_create_product_with_empty_SKU
=== RUN   TestProduct/cannot_create_product_with_invalid_price
=== RUN   TestProduct/update_product_price
=== RUN   TestProduct/cannot_update_price_of_discontinued_product
=== RUN   TestProduct/activate_product
=== RUN   TestProduct/cannot_activate_product_without_stock
=== RUN   TestProduct/cannot_activate_discontinued_product
=== RUN   TestProduct/deactivate_product
=== RUN   TestProduct/cannot_deactivate_inactive_product
=== RUN   TestProduct/discontinue_product
=== RUN   TestProduct/cannot_discontinue_product_with_reserved_stock
=== RUN   TestProduct/add_stock_to_product
=== RUN   TestProduct/reserve_stock
=== RUN   TestProduct/cannot_reserve_stock_from_inactive_product
=== RUN   TestProduct/product_goes_out_of_stock_after_reservation
=== RUN   TestProduct/release_stock
=== RUN   TestProduct/fulfill_stock
=== RUN   TestProduct/check_if_product_is_available_for_order
=== RUN   TestProduct/check_if_product_can_be_deleted
=== RUN   TestProduct/check_low_stock
=== RUN   TestProduct/update_description
=== RUN   TestProduct/update_category
--- PASS: TestProduct (0.00s)
    --- PASS: TestProduct/create_valid_product (0.00s)
    --- PASS: TestProduct/cannot_create_product_with_empty_name (0.00s)
    --- PASS: TestProduct/cannot_create_product_with_empty_SKU (0.00s)
    --- PASS: TestProduct/cannot_create_product_with_invalid_price (0.00s)
    --- PASS: TestProduct/update_product_price (0.00s)
    --- PASS: TestProduct/cannot_update_price_of_discontinued_product (0.00s)
    --- PASS: TestProduct/activate_product (0.00s)
    --- PASS: TestProduct/cannot_activate_product_without_stock (0.00s)
    --- PASS: TestProduct/cannot_activate_discontinued_product (0.00s)
    --- PASS: TestProduct/deactivate_product (0.00s)
    --- PASS: TestProduct/cannot_deactivate_inactive_product (0.00s)
    --- PASS: TestProduct/discontinue_product (0.00s)
    --- PASS: TestProduct/cannot_discontinue_product_with_reserved_stock (0.00s)
    --- PASS: TestProduct/add_stock_to_product (0.00s)
    --- PASS: TestProduct/reserve_stock (0.00s)
    --- PASS: TestProduct/cannot_reserve_stock_from_inactive_product (0.00s)
    --- PASS: TestProduct/product_goes_out_of_stock_after_reservation (0.00s)
    --- PASS: TestProduct/release_stock (0.00s)
    --- PASS: TestProduct/fulfill_stock (0.00s)
    --- PASS: TestProduct/check_if_product_is_available_for_order (0.00s)
    --- PASS: TestProduct/check_if_product_can_be_deleted (0.00s)
    --- PASS: TestProduct/check_low_stock (0.00s)
    --- PASS: TestProduct/update_description (0.00s)
    --- PASS: TestProduct/update_category (0.00s)
PASS

## Application Layer Test Results
=== RUN   TestAddShippingAddressUseCase
=== RUN   TestAddShippingAddressUseCase/successfully_add_shipping_address
=== RUN   TestAddShippingAddressUseCase/customer_not_found
=== RUN   TestAddShippingAddressUseCase/cannot_add_address_with_empty_address_line_1
=== RUN   TestAddShippingAddressUseCase/cannot_add_address_with_invalid_country_code
=== RUN   TestAddShippingAddressUseCase/add_address_with_all_optional_fields
=== RUN   TestAddShippingAddressUseCase/handle_repository_error_when_finding_customer
=== RUN   TestAddShippingAddressUseCase/handle_repository_error_when_updating_customer
--- PASS: TestAddShippingAddressUseCase (0.00s)
    --- PASS: TestAddShippingAddressUseCase/successfully_add_shipping_address (0.00s)
    --- PASS: TestAddShippingAddressUseCase/customer_not_found (0.00s)
    --- PASS: TestAddShippingAddressUseCase/cannot_add_address_with_empty_address_line_1 (0.00s)
    --- PASS: TestAddShippingAddressUseCase/cannot_add_address_with_invalid_country_code (0.00s)
    --- PASS: TestAddShippingAddressUseCase/add_address_with_all_optional_fields (0.00s)
    --- PASS: TestAddShippingAddressUseCase/handle_repository_error_when_finding_customer (0.00s)
    --- PASS: TestAddShippingAddressUseCase/handle_repository_error_when_updating_customer (0.00s)
=== RUN   TestCustomerService
=== RUN   TestCustomerService/register_customer
=== RUN   TestCustomerService/get_customer
=== RUN   TestCustomerService/get_customer_by_email
=== RUN   TestCustomerService/get_customer_by_email_-_not_found
=== RUN   TestCustomerService/update_customer
=== RUN   TestCustomerService/deactivate_customer
=== RUN   TestCustomerService/activate_customer
=== RUN   TestCustomerService/suspend_customer
=== RUN   TestCustomerService/can_customer_place_order_-_yes
=== RUN   TestCustomerService/can_customer_place_order_-_no_addresses
=== RUN   TestCustomerService/can_customer_place_order_-_inactive
=== RUN   TestCustomerService/can_customer_place_order_-_customer_not_found
=== RUN   TestCustomerService/add_shipping_address
=== RUN   TestCustomerService/handle_repository_errors
--- PASS: TestCustomerService (0.00s)
    --- PASS: TestCustomerService/register_customer (0.00s)
    --- PASS: TestCustomerService/get_customer (0.00s)
    --- PASS: TestCustomerService/get_customer_by_email (0.00s)
    --- PASS: TestCustomerService/get_customer_by_email_-_not_found (0.00s)
    --- PASS: TestCustomerService/update_customer (0.00s)
    --- PASS: TestCustomerService/deactivate_customer (0.00s)
    --- PASS: TestCustomerService/activate_customer (0.00s)
    --- PASS: TestCustomerService/suspend_customer (0.00s)
    --- PASS: TestCustomerService/can_customer_place_order_-_yes (0.00s)
    --- PASS: TestCustomerService/can_customer_place_order_-_no_addresses (0.00s)
    --- PASS: TestCustomerService/can_customer_place_order_-_inactive (0.00s)
    --- PASS: TestCustomerService/can_customer_place_order_-_customer_not_found (0.00s)
    --- PASS: TestCustomerService/add_shipping_address (0.00s)
    --- PASS: TestCustomerService/handle_repository_errors (0.00s)
=== RUN   TestGetCustomerUseCase
=== RUN   TestGetCustomerUseCase/successfully_get_customer_by_ID
=== RUN   TestGetCustomerUseCase/customer_not_found
=== RUN   TestGetCustomerUseCase/handle_repository_error
=== RUN   TestGetCustomerUseCase/get_customer_with_no_shipping_addresses
=== RUN   TestGetCustomerUseCase/get_customer_with_multiple_shipping_addresses
--- PASS: TestGetCustomerUseCase (0.00s)
    --- PASS: TestGetCustomerUseCase/successfully_get_customer_by_ID (0.00s)
    --- PASS: TestGetCustomerUseCase/customer_not_found (0.00s)
    --- PASS: TestGetCustomerUseCase/handle_repository_error (0.00s)
    --- PASS: TestGetCustomerUseCase/get_customer_with_no_shipping_addresses (0.00s)
    --- PASS: TestGetCustomerUseCase/get_customer_with_multiple_shipping_addresses (0.00s)
=== RUN   TestUpdateShippingAddressUseCase
=== RUN   TestUpdateShippingAddressUseCase/successfully_update_shipping_address
=== RUN   TestUpdateShippingAddressUseCase/customer_not_found
=== RUN   TestUpdateShippingAddressUseCase/address_not_found
--- PASS: TestUpdateShippingAddressUseCase (0.00s)
    --- PASS: TestUpdateShippingAddressUseCase/successfully_update_shipping_address (0.00s)
    --- PASS: TestUpdateShippingAddressUseCase/customer_not_found (0.00s)
    --- PASS: TestUpdateShippingAddressUseCase/address_not_found (0.00s)
=== RUN   TestRemoveShippingAddressUseCase
=== RUN   TestRemoveShippingAddressUseCase/successfully_remove_shipping_address
=== RUN   TestRemoveShippingAddressUseCase/cannot_remove_only_shipping_address
--- PASS: TestRemoveShippingAddressUseCase (0.00s)
    --- PASS: TestRemoveShippingAddressUseCase/successfully_remove_shipping_address (0.00s)
    --- PASS: TestRemoveShippingAddressUseCase/cannot_remove_only_shipping_address (0.00s)
=== RUN   TestSetDefaultShippingAddressUseCase
=== RUN   TestSetDefaultShippingAddressUseCase/successfully_set_default_shipping_address
=== RUN   TestSetDefaultShippingAddressUseCase/customer_not_found
=== RUN   TestSetDefaultShippingAddressUseCase/address_not_found
--- PASS: TestSetDefaultShippingAddressUseCase (0.00s)
    --- PASS: TestSetDefaultShippingAddressUseCase/successfully_set_default_shipping_address (0.00s)
    --- PASS: TestSetDefaultShippingAddressUseCase/customer_not_found (0.00s)
    --- PASS: TestSetDefaultShippingAddressUseCase/address_not_found (0.00s)
=== RUN   TestRegisterCustomerUseCase
=== RUN   TestRegisterCustomerUseCase/successfully_register_new_customer
=== RUN   TestRegisterCustomerUseCase/cannot_register_customer_with_existing_email
=== RUN   TestRegisterCustomerUseCase/cannot_register_customer_with_invalid_email
=== RUN   TestRegisterCustomerUseCase/cannot_register_customer_with_empty_first_name
=== RUN   TestRegisterCustomerUseCase/cannot_register_customer_with_invalid_phone
=== RUN   TestRegisterCustomerUseCase/handle_repository_error_when_checking_email_existence
=== RUN   TestRegisterCustomerUseCase/handle_repository_error_when_saving_customer
=== RUN   TestRegisterCustomerUseCase/register_customer_without_phone
--- PASS: TestRegisterCustomerUseCase (0.00s)
    --- PASS: TestRegisterCustomerUseCase/successfully_register_new_customer (0.00s)
    --- PASS: TestRegisterCustomerUseCase/cannot_register_customer_with_existing_email (0.00s)
    --- PASS: TestRegisterCustomerUseCase/cannot_register_customer_with_invalid_email (0.00s)
    --- PASS: TestRegisterCustomerUseCase/cannot_register_customer_with_empty_first_name (0.00s)
    --- PASS: TestRegisterCustomerUseCase/cannot_register_customer_with_invalid_phone (0.00s)
    --- PASS: TestRegisterCustomerUseCase/handle_repository_error_when_checking_email_existence (0.00s)
    --- PASS: TestRegisterCustomerUseCase/handle_repository_error_when_saving_customer (0.00s)
    --- PASS: TestRegisterCustomerUseCase/register_customer_without_phone (0.00s)
=== RUN   TestUpdateCustomerUseCase
=== RUN   TestUpdateCustomerUseCase/successfully_update_customer
=== RUN   TestUpdateCustomerUseCase/customer_not_found
=== RUN   TestUpdateCustomerUseCase/cannot_update_with_empty_first_name
=== RUN   TestUpdateCustomerUseCase/cannot_update_with_empty_last_name
=== RUN   TestUpdateCustomerUseCase/cannot_update_with_invalid_phone
=== RUN   TestUpdateCustomerUseCase/update_customer_without_phone
=== RUN   TestUpdateCustomerUseCase/handle_repository_error_when_finding_customer
=== RUN   TestUpdateCustomerUseCase/handle_repository_error_when_updating_customer
--- PASS: TestUpdateCustomerUseCase (0.00s)
    --- PASS: TestUpdateCustomerUseCase/successfully_update_customer (0.00s)
    --- PASS: TestUpdateCustomerUseCase/customer_not_found (0.00s)
    --- PASS: TestUpdateCustomerUseCase/cannot_update_with_empty_first_name (0.00s)
    --- PASS: TestUpdateCustomerUseCase/cannot_update_with_empty_last_name (0.00s)
    --- PASS: TestUpdateCustomerUseCase/cannot_update_with_invalid_phone (0.00s)
    --- PASS: TestUpdateCustomerUseCase/update_customer_without_phone (0.00s)
    --- PASS: TestUpdateCustomerUseCase/handle_repository_error_when_finding_customer (0.00s)
    --- PASS: TestUpdateCustomerUseCase/handle_repository_error_when_updating_customer (0.00s)
PASS

## Coverage Summary
ok  	github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer	(cached)	coverage: 92.9% of statements
ok  	github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/order	(cached)	coverage: 72.3% of statements
ok  	github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/payment	(cached)	coverage: 89.9% of statements
ok  	github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/product	(cached)	coverage: 80.5% of statements
ok  	github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/application/customer	(cached)	coverage: 83.2% of statements
