# Task: cli-chatter-007 - Implement message sending and broadcasting

**User Story**: As a user, I want to send messages to channels so others can see them.

**Acceptance Criteria**:
- Given joined channel, When I type message, Then it's sent to all in channel.
- Given message sent, When received, Then displayed with username and timestamp.
- Given private message, When /msg user message, Then sent only to that user.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 004, 005, 006.
- Assumptions: Real-time via polling or goroutines.
- Risks: Message loss if connection drops.

**Test Notes & Observability Hooks**:
- Test: Send message, check all clients receive.
- Observability: Log messages sent.

**Effort**: M (~6 hours, including broadcasting logic)

**Owners**: @engineer

**PRD References**: #5 M