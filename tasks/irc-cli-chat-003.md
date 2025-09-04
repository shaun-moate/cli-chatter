# Task: Implement Sending Messages

## User Story
As a user, I want to send messages to the chat server so that others can receive my input in real-time.

## Acceptance Criteria
Given an established connection,
When I type a message and press enter,
Then the message is sent to the server and appears in the chat.

## Dependencies / Assumptions / Risks
- Assumption: Connection is already established (depends on task 002).
- Risk: Message loss; mitigation: Implement acknowledgments if possible.

## Test Notes & Observability Hooks
- Manual test: Send message and verify on server side.
- Unit test: Test message formatting and sending logic.
- Observability: Log sent messages.

## Effort
M (Message handling, estimated 8-12 hours)

## Owners
@engineer

## PRD Reference
#key-features, #must-have