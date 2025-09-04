# Task: cli-chatter-012 - Add error logging to stderr

**User Story**: As a developer, I want errors logged to stderr so I can debug issues.

**Acceptance Criteria**:
- Given error occurs, When in any operation, Then logged to stderr with context.
- Given no errors, When running, Then no stderr output.
- Given connection error, When logged, Then includes timestamp.

**Dependencies / Assumptions / Risks**:
- Dependencies: All previous.
- Assumptions: Use fmt.Fprintf(os.Stderr).
- Risks: None.

**Test Notes & Observability Hooks**:
- Test: Cause error, check stderr.
- Observability: This is the hook.

**Effort**: S (~1 hour)

**Owners**: @engineer

**PRD References**: #10
