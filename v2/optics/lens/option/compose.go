package option

// Compose composes two lenses that both return optional values.
//
// This handles the case where both the intermediate structure A and the inner focus B are optional.
// The getter returns None[B] if either A or B is None. The setter behavior is:
//   - Set(Some[B]) when A exists: Updates B in A
//   - Set(Some[B]) when A doesn't exist: Creates A with defaultA and sets B
//   - Set(None[B]) when A doesn't exist: Identity operation (no change)
//   - Set(None[B]) when A exists: Removes B from A (sets it to None)
//
// Type Parameters:
//   - S: Outer structure type
//   - B: Inner focus type (optional)
//   - A: Intermediate structure type (optional)
//
// Parameters:
//   - defaultA: Default value for A when it doesn't exist but B needs to be set
//
// Returns:
//   - A function that takes a LensO[A, B] and returns a function that takes a
//     LensO[S, A] and returns a LensO[S, B]
//
// Example:
//
//	type Settings struct {
//	    MaxRetries *int
//	}
//
//	type Config struct {
//	    Settings *Settings
//	}
//
//	settingsLens := lens.FromNillable(lens.MakeLens(
//	    func(c Config) *Settings { return c.Settings },
//	    func(c Config, s *Settings) Config { c.Settings = s; return c },
//	))
//
//	retriesLens := lens.FromNillable(lens.MakeLensRef(
//	    func(s *Settings) *int { return s.MaxRetries },
//	    func(s *Settings, r *int) *Settings { s.MaxRetries = r; return s },
//	))
//
//	defaultSettings := &Settings{}
//	configRetriesLens := F.Pipe1(settingsLens,
//	    lens.Compose[Config, *int](defaultSettings)(retriesLens))
func Compose[S, B, A any](defaultA A) func(LensO[A, B]) Operator[S, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// Pre-compute setter for Some[A]

// optB is None

// optB is Some

// ComposeOption composes a lens returning an optional value with a lens returning a definite value.
//
// This is useful when you have an optional intermediate structure and want to focus on a field
// within it. The getter returns Option[B] because the container A might not exist. The setter
// behavior depends on the input:
//   - Set(Some[B]): Updates B in A, creating A with defaultA if it doesn't exist
//   - Set(None[B]): Removes A entirely (sets it to None[A])
//
// Type Parameters:
//   - S: Outer structure type
//   - B: Inner focus type (definite value)
//   - A: Intermediate structure type (optional)
//
// Parameters:
//   - defaultA: Default value for A when it doesn't exist but B needs to be set
//
// Returns:
//   - A function that takes a Lens[A, B] and returns a function that takes a
//     LensO[S, A] and returns a LensO[S, B]
//
// Example:
//
//	type Database struct {
//	    Host string
//	    Port int
//	}
//
//	type Config struct {
//	    Database *Database
//	}
//
//	dbLens := lens.FromNillable(lens.MakeLens(
//	    func(c Config) *Database { return c.Database },
//	    func(c Config, db *Database) Config { c.Database = db; return c },
//	))
//
//	portLens := lens.MakeLensRef(
//	    func(db *Database) int { return db.Port },
//	    func(db *Database, port int) *Database { db.Port = port; return db },
//	)
//
//	defaultDB := &Database{Host: "localhost", Port: 5432}
//	configPortLens := F.Pipe1(dbLens, lens.ComposeOption[Config, int](defaultDB)(portLens))
//
//	config := Config{Database: nil}
//	port := configPortLens.Get(config)  // None[int]
//	updated := configPortLens.Set(O.Some(3306))(config)
//	// updated.Database.Port == 3306, Host == "localhost" (from default)
func ComposeOption[S, B, A any](defaultA A) func(Lens[A, B]) Operator[S, A, B] {
	_ = "STUB: not implemented"
	return nil
}

// Pre-compute setters

// optB is None - remove A entirely

// optB is Some - set B
