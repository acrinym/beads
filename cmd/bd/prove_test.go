package main

import "testing"

func TestNormalizeProveTitleStripsPrefixAndPriority(t *testing.T) {
	got := normalizeProveTitle("PhoenixVisualizer-64d P1 Studio audit catalog", `^phoenixvisualizer-[a-z0-9.]+\s*`)
	if got != "studio audit catalog" {
		t.Fatalf("normalizeProveTitle = %q, want %q", got, "studio audit catalog")
	}
}

func TestDuplicateScorePrefersNearIdenticalTitles(t *testing.T) {
	score := duplicateScore("PV window god-class decomposition", "PV window god class decomposition", "")
	if score <= 0.95 {
		t.Fatalf("duplicateScore = %.3f, want > 0.95", score)
	}
}

func TestBuildRecommendationForLikelyCompleted(t *testing.T) {
	rec := buildRecommendation("likely_completed", "open", 0.91, nil, []string{"Open issue has strong completion language in notes/comments."})
	if rec.Action != "review_for_close" {
		t.Fatalf("action = %q, want review_for_close", rec.Action)
	}
	if rec.Confidence != 0.91 {
		t.Fatalf("confidence = %.2f, want 0.91", rec.Confidence)
	}
}
