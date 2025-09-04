# Task: Add Support for Multiple Channels

## User Story
As a user, I want to join and switch between multiple channels so that I can participate in different discussions.

## Acceptance Criteria
Given connected to server,
When I use /join command with channel name,
Then I join the channel and can send/receive messages there.
When I use /switch command,
Then the active channel changes.

## Dependencies / Assumptions / Risks
- Assumption: Server supports channels.
- Risk: State management complexity; mitigation: Use maps for channels.

## Test Notes & Observability Hooks
- Manual test: Join multiple channels.
- Unit test: Test channel switching logic.
- Observability: Log channel joins and switches.

## Effort
M (State management, estimated 12-20 hours)

## Owners
@engineer

## PRD Reference
#key-features, #should-have