-- +goose Up
-- +goose StatementBegin
create index if not exists idx_user_views_viewer_viewed on user_views (viewer_id, viewed_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_user_views_viewer_viewed;
-- +goose StatementEnd
