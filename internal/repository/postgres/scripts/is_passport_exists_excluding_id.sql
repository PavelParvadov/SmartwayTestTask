SELECT EXISTS(
    SELECT 1 FROM passports WHERE type = $1 AND number = $2 AND id <> $3
);



