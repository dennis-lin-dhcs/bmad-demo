---
stepsCompleted: ['step-01-init', 'step-02-discovery', 'step-02b-vision', 'step-02c-executive-summary', 'step-03-success', 'step-04-journeys', 'step-05-domain', 'step-06-innovation', 'step-07-project-type', 'step-08-scoping', 'step-09-functional', 'step-10-nonfunctional', 'step-11-polish', 'step-12-complete']
inputDocuments: []
workflowType: 'prd'
classification:
  projectType: web_app
  domain: healthcare
  complexity: high
  projectContext: greenfield
---

# Product Requirements Document — Medical Forms Platform

**Author:** Dennislin
**Date:** 2026-04-18
**Project:** bmad-demo | Web Application | Healthcare | Greenfield | High Complexity

## Executive Summary

A web-based medical forms platform that digitizes paper-based patient record collection for a single healthcare organization. The system captures and retrieves both standardized clinical forms and custom organizational forms, serving five user roles: patients (form submission), healthcare providers (review, approval, annotation), customer service representatives (annotation), auditors (flagging, annotation), and data scientists (read-only analysis, Phase 2). Multi-tenancy is explicitly out of scope.

The platform eliminates a paper-dependent care continuity workflow — any provider can access complete patient history and resume care without locating a specific colleague. It simultaneously establishes a defensible compliance record base with structured flagging, annotation, and immutable audit trail capabilities as the organization scales.

**Core insight:** The same captured response serves two fundamentally different organizational needs simultaneously — clinical continuity and compliance confidence. The multi-role access model is purpose-built around this dual value; different stakeholders interact with the same data in entirely different ways without interference. Single-org deployment keeps the data model clean and the compliance posture tractable.

## Success Criteria

### User Success

- **Healthcare Providers:** Any provider retrieves a complete patient form history in ≤5 seconds and resumes care without contacting a colleague.
- **Patients:** Form submission completes without confusion; submission confirmation is immediate and unambiguous.
- **Auditors:** Every flagged response is traceable, annotated, and exportable; audit preparation time is measurably reduced vs. paper baseline.
- **Customer Service Reps:** Annotations are addable to any response without altering the clinical record.

### Business Success

- **6-Month Target:** Platform is the single authoritative interface for patient care records — paper retired as primary medium.
- **HIPAA Compliance:** Full HIPAA and California CMIA compliance achieved by end of month 6.

### Technical Success

- Record retrieval: ≤5 seconds for any patient's form history.
- Zero data loss on form submission.
- RBAC enforced for all user roles with no permission bleed.
- Full immutable audit trail on all record access, annotations, flags, and approvals.
- 99.9% uptime; zero unauthorized PHI access incidents post-launch.

## Product Scope

### MVP Strategy

**Approach:** Problem-solving MVP — eliminate the paper-based care continuity gap and establish a defensible compliance record base. No admin UI, no data science access in Phase 1. Prove the core loop before adding operational tooling.

**Team:** 5 engineers. HIPAA compliance overhead and RBAC complexity consume significant capacity on a lean feature set.

### Phase 1 — MVP (Months 0–6)

**User journeys supported:** Patient form submission, provider record retrieval and annotation, CS rep annotation, auditor flagging and annotation.

**Must-have capabilities:**
- Capture and storage of standardized and custom medical forms
- Four-role access: patients, providers, CS reps, auditors
- Provider review and approval workflow
- Annotation for providers, CS reps, and auditors
- Flagging for auditors
- Immutable audit trail on all record interactions
- HIPAA-compliant PHI handling (encryption at rest and in transit)
- WCAG 2.1 AA compliance on patient-facing forms
- Session timeout and role-gated route guards
- Manual user provisioning (IT configures directly — no admin UI)
- Manual form template management (configured at deploy time — no template UI)
- HIPAA compliance sign-off by month 6

**Explicitly excluded from Phase 1:**
- Admin UI for user/role management
- Data scientist access and de-identified export
- SSO/identity provider integration

### Phase 2 — Growth (Months 7–12)

- Admin UI: user account management, role assignment, form template publishing
- Data scientist read-only access with HIPAA-compliant de-identified export
- SSO/SAML identity provider integration
- Patient record export on request (HIPAA right of access)

### Phase 3 — Expansion

- Quality of care auditing across practitioners
- Cost improvement analysis on form data patterns
- Automated compliance and audit report generation
- EHR system integration
- Database remote quorum / multi-AZ replication
- AI-assisted flagging for anomalous responses

### Risk Mitigation

**RBAC (highest technical risk):** Spike RBAC architecture in weeks 1–2 before feature development. Define the permission matrix explicitly and validate with the team before implementation.

**HIPAA / no paper fallback:** Data loss is a patient safety event. Zero data loss architecture required from day one — confirmed write-before-acknowledgment on all submissions. Engage compliance consultant by month 2.

**Team of 5:** Manual provisioning and deferred data scientist access buys ~4–6 weeks of capacity. Protect that buffer — resist scope additions before Phase 1 ships.

## User Journeys

### Journey 1: Maria — Healthcare Provider (Continuity of Care)

Maria arrives for her shift covering Dr. Patel's patients — he called in with a family emergency. The first patient, Mr. Okonkwo, is in the waiting room. Maria has never treated him.

She searches by patient name. Within seconds she has his complete form history: the intake questionnaire from six months ago, the two follow-up forms Dr. Patel completed, and an annotation from last week flagging a medication sensitivity. She reads it in the hallway before walking in.

Mr. Okonkwo doesn't re-explain his history. Maria doesn't apologize for not knowing his case. The handoff is invisible — which is exactly how care continuity should feel.

**Capabilities required:** Patient record lookup, chronological form history, provider annotations visible across providers, search response ≤5 seconds.

---

### Journey 2: James — Patient (Form Submission, Interrupted Session)

James, 58, receives a text link to complete a pre-appointment questionnaire on his phone. The form uses plain-language questions with clear section headers.

Halfway through, his phone rings. He sets it down and returns ten minutes later. The form held his progress. He submits and sees: *"Your responses have been received. Dr. Patel's office will review before your Thursday appointment."*

He doesn't call the front desk to confirm. The confirmation was sufficient.

**Capabilities required:** Mobile-responsive form rendering, session persistence (draft saving), plain-language form support, submission confirmation.

---

### Journey 3: Diana — Auditor (Compliance Review)

Two weeks before an external compliance review, Diana pulls a 10% sample of Q1 records, verifies completeness, and documents findings.

She filters by date range and form type, exports the sample list, and works through each record. Three have incomplete responses. She flags all three, annotates each gap, and generates a flagged-record summary. On review day, every flag, annotation, and access timestamp is present — timestamped, attributed, immutable. The external auditor notes the org has a clear audit trail and structured remediation process.

**Capabilities required:** Record filtering by date/form type, flagging with annotation, exportable audit summary, immutable audit trail with timestamps and attribution.

---

### Journey 4: Carlos — Customer Service Rep (Patient Support)

A patient calls, uncertain her form was saved. Carlos pulls up her record — the form is there, submitted 20 minutes ago. He adds an annotation: *"Patient called to confirm submission — confirmed received, no action needed."*

He confirms the timestamp to her and she hangs up reassured. The annotation ensures any subsequent call handler sees this interaction already occurred.

**Capabilities required:** Record lookup by patient identifier, annotation capability for CS reps, annotation history visible to all authorized roles.

---

### Journey 5: Dr. Chen — Data Scientist (Quarterly Analysis) *(Phase 2)*

Dr. Chen filters intake responses by form type and date range, exports a structured de-identified dataset, and loads it into her notebook. No IT data request, no manual CSV cleaning.

**Capabilities required:** Filtered data export, automatic PHI de-identification by role, structured export format. *(Phase 2)*

---

### Journey 6: Admin — System Administrator (User & Form Management) *(Phase 2)*

A new provider joins. The admin creates an account, assigns the Healthcare Provider role, and sets permissions. She uploads a new custom intake form, maps its fields, and publishes it. The form is live within the hour.

**Capabilities required:** User account management, role assignment, form template management (upload, configure, publish). *(Phase 2)*

---

### Journey Requirements Summary

| Capability | Revealed By | Phase |
|---|---|---|
| Patient record lookup + chronological history | Maria, Carlos | 1 |
| Search response ≤5 seconds | Maria | 1 |
| Provider, CS rep, auditor annotations | Maria, Carlos, Diana | 1 |
| Audit filtering by date/form type + export | Diana | 1 |
| Immutable audit trail (timestamped, attributed) | Diana | 1 |
| Flagging capability | Diana | 1 |
| Mobile-responsive form rendering | James | 1 |
| Draft/session persistence | James | 1 |
| Submission confirmation | James | 1 |
| PHI-aware de-identified data export | Dr. Chen | 2 |
| User & role management | Admin | 2 |
| Form template management | Admin | 2 |

## Domain-Specific Requirements

### Compliance & Regulatory

- **HIPAA:** Full compliance required by month 6. Covers PHI handling, access controls, breach notification procedures, and BAA readiness for third-party vendors.
- **California CMIA:** No medical information disclosed without patient authorization except as permitted by law. System enforces disclosure controls and maintains authorization records for any data access outside the care relationship.
- **HIPAA Security Officer Support:** Platform provides infrastructure to establish a HIPAA Security Officer function — access control management, security incident logging, audit log exports, and user activity reporting.
- **Patient Right of Access:** System supports patient record export on request. *(Phase 2)*

### Data Handling

- **Retention:** All form responses retained for a minimum of 10 years. Deletion is role-gated and logged; no patient self-service deletion.
- **PHI Classification:** All form responses presumed to contain PHI. Encryption at rest and in transit required. RBAC enforces HIPAA minimum necessary access standard.
- **Data Scientist Exports:** De-identified or aggregated per HIPAA Safe Harbor or Expert Determination method. *(Phase 2)*

### Integration Requirements

- **Standalone at Launch:** No EHR integration in Phase 1. Platform is the sole system of record for form responses from day one.
- **SSO/Identity Provider:** SAML/Active Directory integration is Phase 2. Phase 1 supports internal credential management with an SSO migration path.

### Risk Mitigations

- **No Paper Fallback:** Data loss is a patient safety and liability event. Zero data loss required — confirmed write-before-acknowledgment on all submissions; submission receipts provided to patients.
- **High Availability:** Future target is remote quorum / multi-AZ replication. Phase 1 architecture must not preclude this migration.
- **Audit Trail Immutability:** All record access, annotations, flags, and administrative actions logged immutably. No user role — including admins — can edit logs.

## Web Application Requirements

### Platform Overview

Single-page application (SPA) serving all user roles through a unified authenticated interface. Patient-facing form experience is mobile-first. Provider, auditor, CS rep, and admin interfaces are desktop-first but tablet-responsive. No public-facing pages; SEO not required.

### Browser & Platform Matrix

| Platform | Target | Notes |
|---|---|---|
| Chrome (desktop) | Primary | Corporate managed devices |
| Edge (desktop) | Primary | Corporate managed devices |
| Safari (desktop) | Supported | macOS users |
| Firefox (desktop) | Supported | General coverage |
| Chrome (Android) | Primary mobile | Patient form submission |
| Safari (iOS) | Primary mobile | Patient form submission |
| React Native Web | Supported | Future mobile app shell via browser renderer |

Minimum: last 2 major releases per browser.

### Responsive Design

- Patient form submission: fully functional at 320px minimum viewport width
- Provider, auditor, CS rep, admin interfaces: desktop-first, responsive to tablet (768px)
- Data scientist export interface: desktop-only acceptable
- No native device APIs required in Phase 1

### Implementation Constraints

- Client-side routing requires role-based route guards — unauthorized routes redirect, never 404
- PHI rendered in browser must not be cached in browser storage (localStorage, sessionStorage, IndexedDB) — session memory only
- Draft auto-save must not depend on a persistent connection — graceful degradation on slow connections required
- Configurable inactivity session timeout with pre-logout user warning required for HIPAA compliance

## Functional Requirements

### Form Management

- **FR1:** Patients can complete and submit standardized clinical forms
- **FR2:** Patients can complete and submit custom organizational forms
- **FR3:** Patients can save a partially completed form and resume it in a later session
- **FR4:** Patients receive submission confirmation upon successful form submission
- **FR5:** Forms render correctly and are fully operable on mobile browsers

### Record Access & Retrieval

- **FR6:** Healthcare providers can search for patient records by patient identifier
- **FR7:** Healthcare providers can view a patient's complete chronological form history
- **FR8:** Healthcare providers can view any individual form response in full detail
- **FR9:** CS reps can look up patient records by patient identifier
- **FR10:** Auditors can search and filter patient records by date range and form type
- **FR11:** Auditors can view the full detail of any individual form response

### Annotations & Flags

- **FR12:** Healthcare providers can add annotations to any form response
- **FR13:** CS reps can add annotations to any form response
- **FR14:** Auditors can add annotations to any form response
- **FR15:** All authorized roles can view the complete annotation history on any form response
- **FR16:** Auditors can flag a form response for follow-up or compliance review
- **FR17:** Auditors can view a list of all flagged responses

### Review & Approval Workflow

- **FR18:** Healthcare providers can mark a form response as reviewed
- **FR19:** Healthcare providers can approve a form response
- **FR20:** Healthcare providers can view the current review and approval status of any form response

### Audit & Compliance

- **FR21:** The system records an immutable log entry for every record access, annotation, flag, and approval action
- **FR22:** Each audit log entry includes timestamp, user identity, action type, and record identifier
- **FR23:** Auditors can export a filtered summary of audit activity and flagged records
- **FR24:** The system enforces a minimum 10-year retention period on all form responses
- **FR25:** Record deletion requires elevated role authorization and is logged
- **FR26:** Designated security officers can access user activity reports and security incident logs
- **FR27:** Designated security officers can export audit logs for external review

### Authentication & Access Control

- **FR28:** Users authenticate with internal credentials before accessing any system function
- **FR29:** The system enforces role-based access — each role accesses only capabilities defined for that role
- **FR30:** Attempts to access unauthorized routes redirect users to an appropriate destination
- **FR31:** The system terminates inactive sessions after a configurable inactivity period, with prior user warning
- **FR32:** All form responses and PHI are inaccessible to unauthenticated requests

### Data Integrity & Privacy

- **FR33:** The system confirms a successful database write before acknowledging form submission to the patient
- **FR34:** All PHI is encrypted at rest and in transit
- **FR35:** PHI is not persisted in client-side browser storage
- **FR36:** The system enforces disclosure controls — access to medical information outside the care relationship requires documented authorization per CMIA

## Non-Functional Requirements

### Performance

All targets must hold under peak concurrent load of 1,000+ simultaneous authenticated users:

- Patient-facing form pages load within 3 seconds on a 4G mobile connection
- Patient form submission acknowledged within 2 seconds of completion
- Provider patient record search returns results within 5 seconds
- Auditor filtered record view loads within 5 seconds
- Data export initiation completes within 10 seconds for standard query ranges

### Security

- All PHI encrypted at rest (AES-256 or equivalent) and in transit (TLS 1.2 minimum)
- PHI not persisted in client-side browser storage (localStorage, sessionStorage, IndexedDB)
- RBAC enforced server-side — client-side restrictions alone are insufficient
- User sessions expire after configurable inactivity period; users warned before forced logout
- All third-party vendors handling PHI must be capable of executing a HIPAA BAA
- Full HIPAA and California CMIA compliance achieved by end of month 6
- Security incidents automatically logged and accessible to designated security officers

### Reliability

- Availability target: 99.9% uptime (≈8.7 hours unplanned downtime per year)
- Planned maintenance windows outside business hours; advance user notification required
- System confirms successful database write before acknowledging form submission — zero data loss under any failure scenario
- Data loss constitutes a patient safety event; zero data loss is a hard requirement with no exceptions

### Scalability

- Supports 1,000+ concurrent authenticated users without performance degradation beyond defined targets
- Data layer architecture must not preclude future migration to remote quorum / multi-AZ replication topology
- Form schema supports standardized and custom form types without architectural changes per new form type

### Accessibility

- All user-facing interfaces conform to WCAG 2.1 Level AA
- Patient-facing form submission is the highest accessibility priority: keyboard navigation, screen reader compatibility, plain-language labels, and sufficient color contrast required
- Accessibility audit completed and passed before HIPAA compliance sign-off at month 6
