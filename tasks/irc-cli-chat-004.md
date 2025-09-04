# Task: Implement Receiving Messages

## User Story
As a user, I want to receive messages from the chat server so that I can see others' input in real-time.

## Acceptance Criteria
Given an established connection,
When a message is sent from the server,
Then it is displayed in the CLI interface immediately.

## Dependencies / Assumptions / Risks
- Assumption: Connection is established.
- Risk: High latency; mitigation: Optimize reading loop.

## Test Notes & Observability Hooks
- Manual test: Have another client send message.
- Unit test: Simulate incoming messages.
- Observability: Log received messages and timestamps.

## Effort
M (Similar to sending, estimated 8-12 hours)

## Owners
@engineer

## PRD Reference
#key-features, #must-have