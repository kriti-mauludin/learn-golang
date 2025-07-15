CREATE TABLE `special_academy`.`todo_list_categories` 
(`id` INT(100) NOT NULL AUTO_INCREMENT , `name` VARCHAR(100) NOT NULL , `description` TEXT NOT NULL , `created_at` TIMESTAMP NOT NULL , `created_by` INT(100) NOT NULL , PRIMARY KEY (`id`)) 
ENGINE = InnoDB;