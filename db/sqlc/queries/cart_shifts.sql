-- name: CreateCartShift :one
INSERT INTO cart_shifts (id, cart_id, shift_day, start_time, end_time)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetCartShiftsByDay :many
SELECT * FROM cart_shifts WHERE shift_day = $1;

-- name: GetAvailableCartShifts :many
WITH input_params AS (
    SELECT $1::BIGINT AS cart_id, $2::weekday_enum AS shift_day
),
occupied_shifts AS (
    SELECT start_time, end_time 
    FROM cart_shifts 
    WHERE cart_id = (SELECT cart_id FROM input_params)
    AND shift_day = (SELECT shift_day FROM input_params)
)
SELECT time_slot
FROM generate_series('07:00:00'::TIME, '21:00:00'::TIME, '15 minutes'::INTERVAL) AS time_slot
WHERE NOT EXISTS (
    SELECT 1 FROM occupied_shifts
    WHERE time_slot >= occupied_shifts.start_time AND time_slot < occupied_shifts.end_time
)
ORDER BY time_slot;





-- name: UpdateCartShift :exec
UPDATE cart_shifts SET start_time = $3, end_time = $2 WHERE id = $1;

-- name: DeleteCartShift :exec
DELETE FROM cart_shifts WHERE id = $1;
