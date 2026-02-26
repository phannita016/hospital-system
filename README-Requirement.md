# Hospital Middleware System Requirements

## 1. Project Overview
Develop APIs for a Hospital Middleware system to search and display patient information from Hospital Information Systems (HIS).

---

## 2. External HIS Integration
### Hospital A API
- **Route:** `GET https://hospital-a.api.co.th/patient/search/{id}`
- **Request Parameters:**
    - `id` (string): National ID or Passport ID
- **Response Fields (JSON):**
    - `first_name_th`, `middle_name_th`, `last_name_th`
    - `first_name_en`, `middle_name_en`, `last_name_en`
    - `date_of_birth`
    - `patient_hn`
    - `national_id`, `passport_id`
    - `phone_number`, `email`
    - `gender` (M, F)

---

## 3. Database Schema Design
1. **Patient Model**: Design a database schema compatible with hospital data structures.
2. **Staff Model**: Design a database schema where each staff member is associated with a specific hospital. 
    > [!IMPORTANT]
    > Access Control: Staff can only search for patients within the same hospital as themselves.

---

## 4. API Implementation Requirements

### Staff Management
- **Create Staff Member** (`POST /staff/create`)
    - **Input:** `username`, `password`, `hospital`
- **Staff Login** (`POST /staff/login`)
    - **Input:** `username`, `password`, `hospital`

### Patient Management
- **Search Patient** (`GET /patient/search`)
    - **Security:** Requires authentication (Staff login).
    - **Access Control:** Results must only include patients belonging to the same hospital as the staff member.
    - **Search Criteria (Optional):**
        - `national_id`, `passport_id`
        - `first_name`, `middle_name`, `last_name`
        - `date_of_birth`
        - `phone_number`, `email`
    - **Output:** List of matching patients belonging to the same hospital.

---
> [!NOTE]
> All search fields are optional.