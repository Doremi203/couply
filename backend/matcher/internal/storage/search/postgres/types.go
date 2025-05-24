package postgres

const (
	filtersTableName         = "filters"
	filterInterestsTableName = "filter_interests"
	userViewsColumnName      = "user_views"

	userIdColumnName         = "user_id"
	genderPriorityColumnName = "gender_priority"
	minAgeColumnName         = "min_age"
	maxAgeColumnName         = "max_age"
	minHeightColumnName      = "min_height"
	maxHeightColumnName      = "max_height"
	minDistanceKMColumnName  = "min_distance_km"
	maxDistanceKMColumnName  = "max_distance_km"
	goalColumnName           = "goal"
	zodiacColumnName         = "zodiac"
	educationColumnName      = "education"
	childrenColumnName       = "children"
	alcoholColumnName        = "alcohol"
	smokingColumnName        = "smoking"
	onlyVerifiedColumnName   = "only_verified"
	onlyPremiumColumnName    = "only_premium"
	createdAtColumnName      = "created_at"
	updatedAtColumnName      = "updated_at"

	typeColumnName  = "type"
	valueColumnName = "value"

	viewerIDColumnName = "viewer_id"
	viewedIDColumnName = "viewed_id"
)

var (
	filtersColumns = []string{userIdColumnName, genderPriorityColumnName, minAgeColumnName, maxAgeColumnName,
		minHeightColumnName, maxHeightColumnName, minDistanceKMColumnName, maxDistanceKMColumnName, goalColumnName,
		zodiacColumnName, educationColumnName, childrenColumnName, alcoholColumnName, smokingColumnName,
		onlyVerifiedColumnName, onlyPremiumColumnName, createdAtColumnName, updatedAtColumnName}

	filterInterestsColumns = []string{userIdColumnName, typeColumnName, valueColumnName}
)
