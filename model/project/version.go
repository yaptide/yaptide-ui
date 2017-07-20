package project

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// VersionStatus indicate current version status.
type VersionStatus int

const (
	// New version status. It is set during creation or copy/from operation
	// on newly created version.
	New VersionStatus = iota

	// Edited version status. It is set during version modifing.
	Edited

	// Running version status.It is set after simulation/run request
	Running

	// Success version status. It is set after successful simulation.
	Success

	// Failure version status. It is set after unsuccessful simulation
	// associated with simulation engine error.
	Failure

	// Interrupted version status. It is set, when simulation processing is interrupted
	// due to technical difficultes like broken connection or server crash.
	Interrupted
)

// IsModifable return true, if simulation has no started yet,
// so Version content can be changed.
func (v VersionStatus) IsModifable() bool {
	return v == New || v == Edited
}

// IsRunnable return true, if simulation can be runned.
func (v VersionStatus) IsRunnable() bool {
	return v.IsModifable() || v == Interrupted
}

var mapVersionStatusToJSON = map[VersionStatus]string{
	New:         "new",
	Edited:      "edited",
	Running:     "running",
	Success:     "success",
	Failure:     "failure",
	Interrupted: "interrupted",
}

var mapJSONToVersionStatus = map[string]VersionStatus{
	"new":         New,
	"edited":      Edited,
	"running":     Running,
	"success":     Success,
	"failure":     Failure,
	"interrupted": Interrupted,
}

// String fmt.Stringer implementation.
func (v VersionStatus) String() string {
	return mapVersionStatusToJSON[v]
}

// MarshalJSON json.Marshaller implementation.
func (v VersionStatus) MarshalJSON() ([]byte, error) {
	res, ok := mapVersionStatusToJSON[v]
	if !ok {
		return nil,
			fmt.Errorf("VersionStatus.MarshalJSON: can not convert %v to string", v)
	}
	return json.Marshal(res)
}

// UnmarshalJSON json.Unmarshaller implementation.
func (v *VersionStatus) UnmarshalJSON(b []byte) error {
	var input string
	err := json.Unmarshal(b, &input)
	if err != nil {
		return err
	}

	newVersion, ok := mapJSONToVersionStatus[input]
	log.Println(newVersion, ok)
	if !ok {
		return fmt.Errorf(
			"VersionStatus.UnmarshalJSON: can not convert %s to VersionStatus", input)
	}
	*v = newVersion
	return nil
}

// Settings contains simulation settings such as the used simulation engine.
type Settings struct {
	SimulationEngine string `json:"simulationEngine"`
}

// NewSettings constructor.
func NewSettings() Settings {
	return Settings{}
}

// VersionID is key type in Project.Versions list.
type VersionID int

// Version is project version, which contains settting and simulation setup/results.
type Version struct {
	ID       VersionID     `json:"id" bson:"id"`
	Status   VersionStatus `json:"status" bson:"status"`
	Settings Settings      `json:"settings" bson:"settings"`
	SetupID  bson.ObjectId `json:"-" bson:"setup_id"`
	ResultID bson.ObjectId `json:"-" bson:"result_id"`
}
