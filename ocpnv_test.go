package main
import (
    "testing"
    "strings"
    "os"
)

func TestExtract(t *testing.T) {
 want := "BEGIN CERTIFICATE"
 got := Extract("test/test_certs.zip")

 if !strings.Contains(got, want) {
  t.Fatalf("must contain %s, got %s\n", want, got)
 }
}

func TestBuildYML(t *testing.T) {
 want := "machineconfiguration.openshift.io/v1"
 got := BuildYML("test/test_certs.zip")

 if !strings.Contains(got, want) {
  t.Fatalf("must contain %s, got %s\n", want, got)
 }
}

func TestUnzip(t *testing.T) {
 want := "[output-folder/consumer_export.zip]"
 got, _ := Unzip("test/test_certs.zip", "output-folder")
 if _, err := os.Stat(want); err == nil {
  t.Fatalf("expected file %s to exist, got %s\n", want, got)
 }
}
