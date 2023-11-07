# UpdateTaskTasksResponse422ResponseBody

Invalid data posted


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | Raw HTTP response; suitable for custom response parsing                 |                                                                         |
| `Errors`                                                                | [][sdkerrors.UpdateTaskErrors](../../models/errors/updatetaskerrors.md) | :heavy_minus_sign:                                                      | N/A                                                                     |                                                                         |
| `Message`                                                               | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     | The given data was invalid.                                             |