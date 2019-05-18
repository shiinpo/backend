CREATE TABLE IF NOT EXISTS user_records (
  id                serial          PRIMARY KEY,
  weight            INTEGER         NOT NULL,
  reps              INTEGER         NOT NULL,
  rpe               INTEGER         NOT NULL,
  date_performed    DATE            NOT NULL,
  exercise_id       INTEGER         REFERENCES exercise(id) ON DELETE CASCADE ON UPDATE CASCADE,
  user_id           INTEGER         REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  UNIQUE(date_performed, exercise_id)
);