# 🚨 **RULE ENFORCEMENT GUIDE**

## 🎯 **Purpose**

This guide ensures that **ALL coding rules are ALWAYS followed** before any implementation begins. No exceptions. No bypasses.

## 🔄 **Master Rule: @master-enforcement.mdc**

**This rule controls ALL other rules.** It creates a 4-phase system that must be completed sequentially.

### **Phase 1: Ticket Validation** ✅
- **Rule:** @ticket-required.mdc
- **Requirement:** Valid Linear ticket (SMA-XX or ENG-XX format)
- **Status:** BLOCKING - Implementation cannot proceed without this

### **Phase 2: Branch Creation** ✅
- **Rule:** @branching-prs.mdc
- **Requirement:** Branch created from `features` with proper naming
- **Status:** BLOCKING - Implementation cannot proceed without this

### **Phase 3: Quality Standards** ✅
- **Rule:** @go-microservices.mdc
- **Requirement:** All Go best practices planned and documented
- **Status:** BLOCKING - Implementation cannot proceed without this

### **Phase 4: Implementation** 🚀
- **Status:** ALLOWED - Only after all phases complete

## 📋 **Daily Workflow (MANDATORY)**

### **Before ANY Implementation:**

1. **Run Rule Validation:**
   ```bash
   make validate-rules TICKET=SMA-XX BRANCH=feature/SMA-XX-description
   ```

2. **Check All Phases:**
   - [ ] Phase 1: Ticket validated
   - [ ] Phase 2: Branch created
   - [ ] Phase 3: Quality standards planned
   - [ ] Phase 4: Ready to implement

3. **Get User Confirmation:**
   - User must confirm each phase completion
   - AI cannot proceed without explicit approval

### **If ANY Phase Fails:**
- **STOP immediately**
- **Fix the issue**
- **Re-run validation**
- **Get user confirmation again**

## 🚫 **What Gets BLOCKED**

### **Without Valid Ticket:**
- ❌ Any code changes
- ❌ Any file modifications
- ❌ Any implementation work
- ❌ Any feature development

### **Without Proper Branch:**
- ❌ Any commits
- ❌ Any code changes
- ❌ Any implementation work

### **Without Quality Standards:**
- ❌ Any implementation
- ❌ Any code changes
- ❌ Any development work

## ✅ **What Gets ALLOWED**

### **Only After ALL Rules Satisfied:**
- ✅ Code implementation
- ✅ File modifications
- ✅ Feature development
- ✅ Testing and deployment

## 🔧 **How to Use This System**

### **Example 1: New Feature Request**

**User Request:** "Add user management feature"

**AI Response:**
```
🚨 IMPLEMENTATION BLOCKED - RULES NOT SATISFIED

**Missing Requirements:**
- No Linear ticket number
- No branch created
- No quality standards planned

**Required Actions:**
1. Create Linear ticket using @linear-ticket-creator-backend.mdc
2. Create branch from features branch
3. Plan quality standards and patterns

**I will wait for ALL requirements to be satisfied before proceeding.**
```

### **Example 2: Valid Implementation Request**

**User Request:** "Implement SMA-30: Add user management feature"

**AI Response:**
```
🔍 RULE VALIDATION CHECK
=========================

📋 Phase 1: Ticket Validation
✅ PASSED: Ticket format valid (SMA-30)

🌿 Phase 2: Branch Validation
✅ PASSED: Branch format valid (feature/SMA-30-user-management)

🏗️ Phase 3: Quality Standards
✅ PASSED: Go microservices patterns defined
✅ PASSED: Testing strategy planned
✅ PASSED: Error handling patterns defined

📊 VALIDATION SUMMARY
=====================
✅ Ticket format valid (SMA-30)
✅ Branch format valid (feature/SMA-30-user-management)
✅ Quality standards planned

🎉 ALL RULES SATISFIED - IMPLEMENTATION APPROVED

**Ticket:** SMA-30
**Branch:** feature/SMA-30-user-management
**Quality Standards:** GoDoc comments, error handling, context propagation, testing
**Status:** Ready to implement

**Proceeding with implementation...**
```

## 🚨 **Enforcement Mechanisms**

### **1. AI Self-Policing**
- AI must check rules before ANY implementation
- AI must report rule violations immediately
- AI must refuse to proceed if rules not satisfied

### **2. User Enforcement**
- User must provide valid ticket numbers
- User must create proper branches
- User must confirm rule compliance

### **3. Automated Validation**
- Use `make validate-rules` before implementation
- Script validates all rule requirements
- Exit codes indicate compliance status

## 📝 **Rule Violation Reporting**

### **If AI Violates Rules:**
1. **Immediate stop** of all work
2. **Report violation** to user
3. **Restart** from Phase 1
4. **No exceptions** allowed

### **If User Bypasses Rules:**
1. **AI must refuse** to proceed
2. **Explain why** rules are mandatory
3. **Offer help** to satisfy requirements
4. **Wait** for proper setup

## 🎯 **Success Metrics**

### **Rule Compliance Rate:**
- **Target:** 100% compliance
- **Measurement:** All implementations follow rules
- **Tracking:** Daily validation checks

### **Quality Improvements:**
- **Target:** All code follows Go best practices
- **Measurement:** Code review compliance
- **Tracking:** Automated validation results

## 🔄 **Integration with Existing Workflows**

### **Linear Integration:**
- All work must have Linear tickets
- Ticket numbers used in branch names
- Ticket context in all commits

### **Git Workflow:**
- All branches from `features`
- Proper naming conventions
- PRs against `features` branch

### **Quality Assurance:**
- GoDoc comments required
- Error handling patterns
- Testing coverage requirements

## 🚨 **CRITICAL REMINDERS**

1. **This system is MANDATORY** - no exceptions
2. **All phases must complete** before implementation
3. **User confirmation required** at each phase
4. **Validation script must pass** before proceeding
5. **Rules cannot be bypassed** or ignored

## 📞 **Getting Help**

### **If Rules Are Confusing:**
1. Review this guide
2. Check individual rule files
3. Run validation script
4. Ask for clarification

### **If Rules Are Too Strict:**
1. Rules are designed for quality
2. They prevent common mistakes
3. They ensure consistency
4. They improve maintainability

## 🎉 **Benefits of Following Rules**

1. **Consistent Quality** - All code follows standards
2. **Better Testing** - Comprehensive test coverage
3. **Easier Maintenance** - Clear patterns and structure
4. **Team Collaboration** - Standardized workflows
5. **Professional Standards** - Industry best practices

---

**Remember: Rules exist to make your code better, not to slow you down.**
**Follow them, and you'll build better software faster.**
