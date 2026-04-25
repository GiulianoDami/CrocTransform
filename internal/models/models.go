package models

import "time"

// FossilData represents the structured data extracted from fossil specimens
type FossilData struct {
	// Unique identifier for the fossil specimen
	ID string `json:"id"`
	
	// Taxonomic classification of the specimen
	Species string `json:"species"`
	
	// Geological age in millions of years
	Age float64 `json:"age"`
	
	// Location where the fossil was discovered
	Location string `json:"location"`
	
	// Morphological measurements of limb bones
	LimbMeasurements map[string]float64 `json:"limb_measurements"`
	
	// Joint angle measurements
	JointAngles map[string]float64 `json:"joint_angles"`
	
	// Growth stage indicators
	GrowthStage string `json:"growth_stage"`
	
	// Collection date
	CollectionDate time.Time `json:"collection_date"`
	
	// Quality assessment score (0-1)
	QualityScore float64 `json:"quality_score"`
}

// AnalysisResult represents the output of the locomotion pattern analysis
type AnalysisResult struct {
	// Reference to the input fossil data
	FossilID string `json:"fossil_id"`
	
	// Detected locomotion pattern
	Pattern string `json:"pattern"`
	
	// Confidence level of the detection (0-1)
	Confidence float64 `json:"confidence"`
	
	// Statistical significance of the pattern
	Significance float64 `json:"significance"`
	
	// Analysis timestamp
	Timestamp time.Time `json:"timestamp"`
	
	// Detailed findings and observations
	Findings []string `json:"findings"`
	
	// Classification of the locomotion transition
	TransitionType string `json:"transition_type"`
	
	// Probability scores for different locomotion categories
	Probabilities map[string]float64 `json:"probabilities"`
}