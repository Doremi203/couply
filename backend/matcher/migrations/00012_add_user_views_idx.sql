-- +goose Up
-- +goose StatementBegin
create index concurrently if not exists idx_user_views_viewer_viewed on user_views (viewer_id, viewed_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently if exists idx_user_views_viewer_viewed;
-- +goose StatementEnd
