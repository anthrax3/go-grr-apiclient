// This file has been autogenerated by ./generate/gen-functions.sh

package apiclient

func (c APIClient) DeleteArtifacts(rq ApiDeleteArtifactsArgs) error {
	return c.post("/api/artifacts/delete", &rq)
}
