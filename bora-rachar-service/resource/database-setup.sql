-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema BORA_RACHAR
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema BORA_RACHAR
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `BORA_RACHAR` DEFAULT CHARACTER SET utf8 ;
USE `BORA_RACHAR` ;

-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`user` (
                                                    `id` INT NOT NULL,
                                                    `user_id` VARCHAR(45) NOT NULL,
                                                    `name` VARCHAR(45) NULL,
                                                    `email` VARCHAR(45) NOT NULL,
                                                    `pix_key` VARCHAR(45) NULL,
                                                    `created_at` DATETIME NULL,
                                                    `updated_at` DATETIME NULL,
                                                    `deleted_at` DATETIME NULL,
                                                    PRIMARY KEY (`id`),
                                                    UNIQUE INDEX `userId_UNIQUE` (`user_id` ASC) VISIBLE)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`group`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`group` (
                                                     `id` INT NOT NULL,
                                                     `name` VARCHAR(45) NOT NULL,
                                                     `avatar` VARCHAR(45) NULL,
                                                     `access_code` VARCHAR(45) NULL,
                                                     `created_by` VARCHAR(45) NULL,
                                                     `created_at` DATETIME NULL,
                                                     `updated_at` DATETIME NULL,
                                                     `deleted_at` DATETIME NULL,
                                                     PRIMARY KEY (`id`),
                                                     UNIQUE INDEX `access_code_UNIQUE` (`access_code` ASC) VISIBLE,
                                                     INDEX `userId_idx` (`created_by` ASC) VISIBLE,
                                                     CONSTRAINT `userId_group`
                                                         FOREIGN KEY (`created_by`)
                                                             REFERENCES `BORA_RACHAR`.`user` (`user_id`)
                                                             ON DELETE NO ACTION
                                                             ON UPDATE NO ACTION)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`expense`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`expense` (
                                                       `id` INT NOT NULL,
                                                       `group_id` INT NULL,
                                                       `title` VARCHAR(45) NULL,
                                                       `description` VARCHAR(45) NULL,
                                                       `category` VARCHAR(45) NULL,
                                                       `avatar` VARCHAR(45) NULL,
                                                       `value` DECIMAL(19,4) NULL,
                                                       `expense_date` DATETIME NULL,
                                                       `created_by` VARCHAR(45) NULL,
                                                       `created_at` DATETIME NULL,
                                                       `updated_at` DATETIME NULL,
                                                       `deleted_at` DATETIME NULL,
                                                       PRIMARY KEY (`id`),
                                                       INDEX `id_grupo.expense_idx` (`group_id` ASC) VISIBLE,
                                                       INDEX `created_by_expense_idx` (`created_by` ASC) VISIBLE,
                                                       CONSTRAINT `id_grupo_expense`
                                                           FOREIGN KEY (`group_id`)
                                                               REFERENCES `BORA_RACHAR`.`group` (`id`)
                                                               ON DELETE NO ACTION
                                                               ON UPDATE NO ACTION,
                                                       CONSTRAINT `created_by_expense`
                                                           FOREIGN KEY (`created_by`)
                                                               REFERENCES `BORA_RACHAR`.`user` (`user_id`)
                                                               ON DELETE NO ACTION
                                                               ON UPDATE NO ACTION)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`group_participant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`group_participant` (
                                                                 `id` INT NOT NULL,
                                                                 `user_id` VARCHAR(45) NULL,
                                                                 `group_id` INT NULL,
                                                                 `is_admin` BLOB NULL,
                                                                 `admission_date` DATETIME NULL,
                                                                 `removed_at` DATETIME NULL,
                                                                 PRIMARY KEY (`id`),
                                                                 INDEX `user_id_group_participant_idx` (`user_id` ASC) VISIBLE,
                                                                 INDEX `group_id_group_participant_idx` (`group_id` ASC) VISIBLE,
                                                                 CONSTRAINT `user_id_group_participant`
                                                                     FOREIGN KEY (`user_id`)
                                                                         REFERENCES `BORA_RACHAR`.`user` (`user_id`)
                                                                         ON DELETE NO ACTION
                                                                         ON UPDATE NO ACTION,
                                                                 CONSTRAINT `group_id_group_participant`
                                                                     FOREIGN KEY (`group_id`)
                                                                         REFERENCES `BORA_RACHAR`.`group` (`id`)
                                                                         ON DELETE NO ACTION
                                                                         ON UPDATE NO ACTION)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`group_solicitation`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`group_solicitation` (
                                                                  `id` INT NOT NULL,
                                                                  `user_id` VARCHAR(45) NULL,
                                                                  `group_id` INT NULL,
                                                                  `created_at` DATETIME NULL,
                                                                  PRIMARY KEY (`id`),
                                                                  INDEX `group_id_idx` (`group_id` ASC) VISIBLE,
                                                                  INDEX `user_id_idx` (`user_id` ASC) VISIBLE,
                                                                  CONSTRAINT `user_id_group_solicitation`
                                                                      FOREIGN KEY (`user_id`)
                                                                          REFERENCES `BORA_RACHAR`.`user` (`user_id`)
                                                                          ON DELETE NO ACTION
                                                                          ON UPDATE NO ACTION,
                                                                  CONSTRAINT `group_id_group_solicitation`
                                                                      FOREIGN KEY (`group_id`)
                                                                          REFERENCES `BORA_RACHAR`.`group` (`id`)
                                                                          ON DELETE NO ACTION
                                                                          ON UPDATE NO ACTION)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`expense_payment_split`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`expense_payment_split` (
                                                                     `id` INT NOT NULL,
                                                                     `user_id` VARCHAR(45) NULL,
                                                                     `expense_id` INT NULL,
                                                                     `value` DECIMAL(19,4) NULL,
                                                                     `transaction_type` ENUM('P', 'B') NULL,
                                                                     `created_at` DATETIME NULL,
                                                                     `is_debt_settled` TINYINT NULL,
                                                                     PRIMARY KEY (`id`),
                                                                     INDEX `expense_id_expense_borrower_idx` (`expense_id` ASC) VISIBLE,
                                                                     INDEX `user_id_expense_borrower_idx` (`user_id` ASC) VISIBLE,
                                                                     CONSTRAINT `expense_id_expense_payment_split`
                                                                         FOREIGN KEY (`expense_id`)
                                                                             REFERENCES `BORA_RACHAR`.`expense` (`id`)
                                                                             ON DELETE NO ACTION
                                                                             ON UPDATE NO ACTION,
                                                                     CONSTRAINT `user_id_expense_payment_split`
                                                                         FOREIGN KEY (`user_id`)
                                                                             REFERENCES `BORA_RACHAR`.`user` (`user_id`)
                                                                             ON DELETE NO ACTION
                                                                             ON UPDATE NO ACTION)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `BORA_RACHAR`.`expense_commentary`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `BORA_RACHAR`.`expense_commentary` (
                                                                  `id` INT NOT NULL,
                                                                  `user_id` VARCHAR(45) NULL,
                                                                  `expense_id` INT NULL,
                                                                  `comentary` LONGTEXT NULL,
                                                                  `created_at` DATETIME NULL,
                                                                  `updated_at` DATETIME NULL,
                                                                  `deleted_at` DATETIME NULL,
                                                                  PRIMARY KEY (`id`),
                                                                  INDEX `expense_id_expense_commentary_idx` (`expense_id` ASC) VISIBLE,
                                                                  INDEX `user_id_expense_commentary_idx` (`user_id` ASC) VISIBLE,
                                                                  CONSTRAINT `expense_id_expense_commentary`
                                                                      FOREIGN KEY (`expense_id`)
                                                                          REFERENCES `BORA_RACHAR`.`expense` (`id`)
                                                                          ON DELETE NO ACTION
                                                                          ON UPDATE NO ACTION,
                                                                  CONSTRAINT `user_id_expense_commentary`
                                                                      FOREIGN KEY (`user_id`)
                                                                          REFERENCES `BORA_RACHAR`.`user` (`user_id`)
                                                                          ON DELETE NO ACTION
                                                                          ON UPDATE NO ACTION)
    ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
