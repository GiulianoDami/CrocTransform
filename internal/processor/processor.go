package processor

// FossilData represents the input data structure for fossil analysis
type FossilData struct {
	// Limb bone proportions (femur/tibia, humerus/ulna ratios)
	LimbProportions map[string]float64
	
	// Joint angles measured from fossil specimens
	JointAngles map[string]float64
	
	// Growth stage indicators (juvenile/adult)
	GrowthStage string
	
	// Body size measurements
	BodyLength float64
	BodyHeight float64
}

// AnalysisResult represents the output of the fossil analysis
type AnalysisResult struct {
	// Confidence score for locomotion transition detection
	Confidence float64
	
	// Detected locomotion pattern
	LocomotionPattern string
	
	// Key findings from the analysis
	KeyFindings []string
	
	// Whether a transition was detected
	TransitionDetected bool
}

// AnalyzeFossil performs core analysis on fossil data to detect locomotion transitions
func AnalyzeFossil(data FossilData) AnalysisResult {
	result := AnalysisResult{
		Confidence:         0.0,
		LocomotionPattern:  "unknown",
		KeyFindings:        []string{},
		TransitionDetected: false,
	}
	
	// Check if we have sufficient data
	if len(data.LimbProportions) == 0 || len(data.JointAngles) == 0 {
		result.KeyFindings = append(result.KeyFindings, "Insufficient morphological data")
		return result
	}
	
	// Analyze limb proportions for transition indicators
	quadrupedalIndicators := 0
	bipedalIndicators := 0
	
	// Check femur/tibia ratio (typically more balanced in bipedal forms)
	if ratio, exists := data.LimbProportions["femur_tibia"]; exists {
		if ratio > 1.2 {
			bipedalIndicators++
		} else {
			quadrupedalIndicators++
		}
	}
	
	// Check humerus/ulna ratio (more variable in bipedal forms)
	if ratio, exists := data.LimbProportions["humerus_ulna"]; exists {
		if ratio > 1.1 {
			bipedalIndicators++
		} else {
			quadrupedalIndicators++
		}
	}
	
	// Analyze joint angles for stance differences
	if angle, exists := data.JointAngles["hip_angle"]; exists {
		if angle < 90 {
			quadrupedalIndicators++
		} else {
			bipedalIndicators++
		}
	}
	
	if angle, exists := data.JointAngles["knee_angle"]; exists {
		if angle > 150 {
			bipedalIndicators++
		} else {
			quadrupedalIndicators++
		}
	}
	
	// Determine locomotion pattern based on indicators
	if quadrupedalIndicators > bipedalIndicators {
		result.LocomotionPattern = "quadrupedal"
	} else if bipedalIndicators > quadrupedalIndicators {
		result.LocomotionPattern = "bipedal"
	} else {
		result.LocomotionPattern = "ambiguous"
	}
	
	// Calculate confidence based on data consistency
	totalIndicators := quadrupedalIndicators + bipedalIndicators
	if totalIndicators > 0 {
		result.Confidence = float64(max(quadrupedalIndicators, bipedalIndicators)) / float64(totalIndicators)
	}
	
	// Add key findings
	if result.LocomotionPattern != "ambiguous" {
		result.KeyFindings = append(result.KeyFindings, 
			"Primary locomotion pattern identified as "+result.LocomotionPattern)
	}
	
	// Check for potential transition signs
	if data.GrowthStage == "juvenile" && result.LocomotionPattern == "quadrupedal" {
		result.KeyFindings = append(result.KeyFindings, 
			"Juvenile specimen with quadrupedal characteristics")
	} else if data.GrowthStage == "adult" && result.LocomotionPattern == "bipedal" {
		result.KeyFindings = append(result.KeyFindings, 
			"Adult specimen with bipedal characteristics")
	}
	
	// Detect possible transition if juvenile shows quadrupedal and adult shows bipedal
	result.TransitionDetected = (quadrupedalIndicators > 0 && bipedalIndicators > 0)
	
	return result
}

// Helper function to find maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}