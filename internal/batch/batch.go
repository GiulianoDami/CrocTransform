package batch

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"croctransform/internal/fossil"
)

// ProcessBatch processes multiple fossil specimens from input directory and writes results to CSV
func ProcessBatch(inputDir, outputCSV string) error {
	// Create output CSV file
	outputFile, err := os.Create(outputCSV)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write header
	header := []string{"filename", "specimen_id", "locomotion_pattern", "confidence_score", "notes"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Walk through input directory
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Process only supported file types
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".json" && ext != ".xml" {
			return nil
		}

		// Parse fossil data
		result, err := fossil.AnalyzeSpecimen(path)
		if err != nil {
			fmt.Printf("Warning: Failed to process %s: %v\n", path, err)
			return nil
		}

		// Write result to CSV
		record := []string{
			filepath.Base(path),
			result.SpecimenID,
			result.LocomotionPattern,
			fmt.Sprintf("%.2f", result.ConfidenceScore),
			result.Notes,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}

	return nil
}