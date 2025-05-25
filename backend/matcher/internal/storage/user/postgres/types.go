package postgres

const (
	usersTableName     = "users"
	photosTableName    = "photos"
	interestsTableName = "interests"

	idColumnName         = "id"
	nameColumnName       = "name"
	ageColumnName        = "age"
	genderColumnName     = "gender"
	latitudeColumnName   = "latitude"
	longitudeColumnName  = "longitude"
	bioColumnName        = "bio"
	goalColumnName       = "goal"
	zodiacColumnName     = "zodiac"
	heightColumnName     = "height"
	educationColumnName  = "education"
	childrenColumnName   = "children"
	alcoholColumnName    = "alcohol"
	smokingColumnName    = "smoking"
	isHiddenColumnName   = "is_hidden"
	isVerifiedColumnName = "is_verified"
	isPremiumColumnName  = "is_premium"
	isBlockedColumnName  = "is_blocked"
	createdAtColumnName  = "created_at"
	updatedAtColumnName  = "updated_at"

	userIDColumnName      = "user_id"
	orderNumberColumnName = "order_number"
	objectKeyColumnName   = "object_key"
	mimeTypeColumnName    = "mime_type"
	uploadedAtColumnName  = "uploaded_at"

	typeColumnName  = "type"
	valueColumnName = "value"
)

var (
	usersColumns = []string{idColumnName, nameColumnName, ageColumnName, genderColumnName, latitudeColumnName,
		longitudeColumnName, bioColumnName, goalColumnName, zodiacColumnName, heightColumnName, educationColumnName,
		childrenColumnName, alcoholColumnName, smokingColumnName, isHiddenColumnName, isVerifiedColumnName, isPremiumColumnName,
		isBlockedColumnName, createdAtColumnName, updatedAtColumnName}

	photosColumns = []string{userIDColumnName, orderNumberColumnName, objectKeyColumnName, mimeTypeColumnName, uploadedAtColumnName}

	interestsColumns = []string{userIDColumnName, typeColumnName, valueColumnName}
)
