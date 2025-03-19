-- name: GetEnumMeetingTypes :many
SELECT unnest(enum_range(NULL::meeting_type_enum)) AS value;

-- name: GetEnumAssignmentTypes :many
SELECT unnest(enum_range(NULL::assignment_type_enum)) AS value;

-- name: GetEnumCleaningTypes :many
SELECT unnest(enum_range(NULL::cleaning_type_enum)) AS value;
