# Task: cli-chatter-013 - Add command usage tracking

**User Story**: As a product manager, I want command usage tracked so I can measure adoption.

**Acceptance Criteria**:
- Given command executed, When /join or /msg, Then usage counted.
- Given metrics requested, When /stats, Then shows counts.
- Given persisted, When restart, Then counts maintained.

**Dependencies / Assumptions / Risks**:
- Dependencies: Tasks 010.
- Assumptions: In-memory counters, persisted.
- Risks: None.

**Test Notes & Observability Hooks**:
- Test: Execute commands, check /stats.
- Observability: Metrics for success metrics.

**Effort**: S (~2 hours)

**Owners**: @engineer

**PRD References**: #4, #10
