# Product Requirements Document: CLI Chatter

**Slug**: cli-chatter  
**Date**: Thu Sep 04 2025  
**Author**: Senior Product Manager  

## 1. Problem Statement
Developers and sysadmins need a lightweight, CLI-based chat application for quick communication in terminal environments. Existing IRC clients are often GUI-heavy or require external servers, while modern chat tools lack CLI integration. This app will provide IRC-inspired functionality (channels, commands) in a pure CLI interface for efficient, distraction-free chatting.

## 2. User Personas
- **Primary**: DevOps engineers who work in terminals and need quick team coordination without switching contexts.
- **Secondary**: Developers debugging in CLI who want to chat without leaving the terminal.
- **Tertiary**: Sysadmins managing servers via SSH who need simple communication tools.

## 3. Goals & Objectives
- Deliver a functional CLI chat app that mimics IRC basics (join channels, send messages, basic commands).
- Ensure it's easy to build/run with `go run main.go`.
- Follow Go best practices from AGENTS.md for maintainable code.

## 4. Success Metrics
- User adoption: 80% of target users can install and run without issues.
- Performance: Handles 10+ concurrent users/channels without lag.
- Usability: Average session time >5 minutes; low error rate (<5%).

## 5. Features (MoSCoW Prioritization)
- **Must Have (M)**: Basic message sending/receiving, channel joining, user commands (/join, /msg).
- **Should Have (S)**: User authentication (simple username), message history.
- **Could Have (C)**: File sharing, encryption.
- **Won't Have (W)**: GUI mode, video calls, integrations with external services.

## 6. Non-Goals
- Not a full IRC server replacement (no federation).
- No mobile/web clients.
- Avoid heavy dependencies; keep it lightweight.

## 7. Risks
- CLI limitations for real-time updates (may need polling).
- Networking complexity if expanding beyond local.
- User adoption if not intuitive.

## 8. Dependencies
- Go 1.22.0 or higher (confirmed on darwin/arm64).
- Potential: net/http for basic networking (if needed for multi-user).

## 9. MVP Slice
Core: Echo loop (existing) + channel support + basic commands. Launch with local file-based persistence.

## 10. Observability
- Log errors to stderr.
- Track command usage for metrics.

## 11. Decision Log
- Initial: Use Go for CLI strength and existing project setup.
- ADR: CLI-only to maintain simplicity.