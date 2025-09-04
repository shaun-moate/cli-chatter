# Task: Implement Private Messaging

## User Story
As a user, I want to send private messages to other users so that I can have one-on-one conversations.

## Acceptance Criteria
Given connected and user list available,
When I use /msg command with username and message,
Then the message is sent privately and received only by that user.

## Dependencies / Assumptions / Risks
- Assumption: Server supports private messages.
- Risk: User discovery; mitigation: Implement /who command.

## Test Notes & Observability Hooks
- Manual test: Send private message.
- Unit test: Test message routing.
- Observability: Log private messages (with privacy considerations).

## Effort
M (Similar to public messages, estimated 8-16 hours)

## Owners
@engineer

## PRD Reference
#key-features, #should-have