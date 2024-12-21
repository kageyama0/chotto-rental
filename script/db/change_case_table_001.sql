-- 一時的に外部キー制約を削除
ALTER TABLE applications DROP CONSTRAINT applications_case_id_fkey;
ALTER TABLE applications DROP CONSTRAINT fk_applications_case;
ALTER TABLE matchings DROP CONSTRAINT matchings_case_id_fkey;
ALTER TABLE matchings DROP CONSTRAINT fk_matchings_case;

-- location列を分割し、新しい列を追加
ALTER TABLE cases
  ADD COLUMN category varchar(255) NOT NULL DEFAULT 'other',
  ADD COLUMN required_people integer NOT NULL DEFAULT 1,
  ADD COLUMN start_time varchar(5) NOT NULL DEFAULT '00:00',
  ADD COLUMN prefecture varchar(255) NOT NULL DEFAULT '',
  ADD COLUMN city varchar(255) NOT NULL DEFAULT '',
  ADD COLUMN address text,
  -- 既存の列の制約を追加
  ALTER COLUMN title TYPE varchar(100),
  ALTER COLUMN description TYPE varchar(2000),

ALTER TABLE cases RENAME COLUMN duration_minutes TO duration;

-- Check制約の追加
ALTER TABLE cases
  ADD CONSTRAINT check_reward CHECK (reward >= 500 AND reward <= 100000),
  ADD CONSTRAINT check_required_people CHECK (required_people >= 1 AND required_people <= 10),
  ADD CONSTRAINT check_duration CHECK (duration >= 15 AND duration <= 360);

-- locationデータの移行のために一時的なトリガーを作成することもできます
-- ここではデフォルト値を設定しているので、既存データは手動で更新する必要があります

-- 既存データの移行 (例として location を分割する場合)
-- UPDATE cases SET
--   prefecture = split_part(location, ' ', 1),
--   city = split_part(location, ' ', 2),
--   address = split_part(location, ' ', 3);

-- 不要になったカラムの削除
ALTER TABLE cases DROP COLUMN location;

-- 外部キー制約を再作成
ALTER TABLE applications
  ADD CONSTRAINT fk_applications_case
  FOREIGN KEY (case_id)
  REFERENCES cases(id);

ALTER TABLE matchings
  ADD CONSTRAINT fk_matchings_case
  FOREIGN KEY (case_id)
  REFERENCES cases(id);

-- インデックスの作成（必要に応じて）
CREATE INDEX idx_cases_category ON cases(category);
CREATE INDEX idx_cases_prefecture ON cases(prefecture);
