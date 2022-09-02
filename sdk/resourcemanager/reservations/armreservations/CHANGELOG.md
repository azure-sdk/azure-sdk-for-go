# Release History

## 1.1.0 (2022-09-02)
### Features Added

- New function `*ReservationClient.Unarchive(context.Context, string, string, *ReservationClientUnarchiveOptions) (ReservationClientUnarchiveResponse, error)`
- New function `*ReservationClient.Archive(context.Context, string, string, *ReservationClientArchiveOptions) (ReservationClientArchiveResponse, error)`
- New struct `ReservationClientArchiveOptions`
- New struct `ReservationClientArchiveResponse`
- New struct `ReservationClientUnarchiveOptions`
- New struct `ReservationClientUnarchiveResponse`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).