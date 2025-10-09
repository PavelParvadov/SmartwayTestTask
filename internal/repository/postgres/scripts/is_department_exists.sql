SELECT EXISTS(
    SELECT 1 FROM departments WHERE id = $1
)