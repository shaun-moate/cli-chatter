# Task: cli-chatter-008 - Implement /join command

**User Story**: As a user, I want to join channels using /join so I can participate in group chats.

**Acceptance Criteria**:
- Given connected, When I type /join #channel, Then I join the channel.
- Given not joined, When I send message, Then it's to current channel.
- Given invalid channel, When /join invalid, Then error message.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 003, 005.
- Assumptions: Channels created on join if not exist.
- Risks: None.

**Test Notes & Observability Hooks**:
- Test: Type /join, check joined.
- Observability: Log command usage.

**Effort**: S (~2 hours)

**Owners**: @engineer

**PRD References**: #5 M
