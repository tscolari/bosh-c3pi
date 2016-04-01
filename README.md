


# Usage:

1. Create a class to talk with your infrastructure that implements the `cloud.Cloud` interface
2. In your `main` function instantiate a `Runner` object with your cloud implementation:

```go
import (
	"github.com/tscolari/bosh-c3pi/cpi"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

  func main() {
    cpiClient := yourimplementation.New()
    logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
    runner := cpi.NewRunner(cpiClient, logger)

    runner.Run(os.Stdin, os.Stderr)
  }
```

3. done

# Further steps:

You still need to wrap it in a bosh release, to make it available to bosh. (WIP)
