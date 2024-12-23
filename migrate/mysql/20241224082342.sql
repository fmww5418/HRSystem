-- Create "organizations" table
CREATE TABLE `organizations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `admin_id` bigint unsigned NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_organizations_admin` (`admin_id`),
  INDEX `idx_organizations_deleted_at` (`deleted_at`),
  UNIQUE INDEX `uni_organizations_name` (`name`),
  CONSTRAINT `fk_organizations_admin` FOREIGN KEY (`admin_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "departments" table
CREATE TABLE `departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `organization_id` bigint unsigned NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_departments_organization` (`organization_id`),
  INDEX `idx_departments_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_departments_organization` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Modify "employees" table
ALTER TABLE `employees` ADD COLUMN `department_id` bigint unsigned NULL, ADD COLUMN `supervisor_employee_id` bigint unsigned NULL, ADD INDEX `fk_employees_department` (`department_id`), ADD INDEX `fk_employees_supervisor_employee` (`supervisor_employee_id`), ADD CONSTRAINT `fk_employees_department` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT `fk_employees_supervisor_employee` FOREIGN KEY (`supervisor_employee_id`) REFERENCES `employees` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
