package pixela

import "testing"

const (
	pixelDate     = "20200101"
	pixelQuantity = "10"
)

func testE2EPixelCreate(t *testing.T) {
	input := &PixelCreateInput{
		Date:         String(pixelDate),
		Quantity:     String("1"),
		OptionalData: String("{\"key\":\"value\"}"),
		GraphID:      String(graphID),
	}
	result, err := e2eClient.Pixel().Create(input)
	if err != nil {
		t.Errorf("Pixel.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Pixel.Create() got: %+v\nwant: true", result)
	}
}

func testE2EPixelIncrement(t *testing.T) {
	input := &PixelIncrementInput{GraphID: String(graphID)}
	result, err := e2eClient.Pixel().Increment(input)
	if err != nil {
		t.Errorf("Pixel.Increment() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Pixel.Increment() got: %+v\nwant: true", result)
	}
}

func testE2EPixelDecrement(t *testing.T) {
	input := &PixelDecrementInput{GraphID: String(graphID)}
	result, err := e2eClient.Pixel().Decrement(input)
	if err != nil {
		t.Errorf("Pixel.Decrement() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Pixel.Decrement() got: %+v\nwant: true", result)
	}
}

func testE2EPixelUpdate(t *testing.T) {
	input := &PixelUpdateInput{
		Date:     String(pixelDate),
		Quantity: String(pixelQuantity),
		GraphID:  String(graphID),
	}
	result, err := e2eClient.Pixel().Update(input)
	if err != nil {
		t.Errorf("Pixel.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Pixel.Update() got: %+v\nwant: true", result)
	}
}

func testE2EPixelGet(t *testing.T) {
	input := &PixelGetInput{
		Date:    String(pixelDate),
		GraphID: String(graphID),
	}
	result, err := e2eClient.Pixel().Get(input)
	if err != nil {
		t.Errorf("Pixel.Get() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Pixel.Get() got: %+v\nwant: true", result)
	}
	if result.Quantity != pixelQuantity {
		t.Errorf("Pixel.Get() got: %s\nwant: %s", result.Quantity, pixelQuantity)
	}
}

func testE2EPixelDelete(t *testing.T) {
	input := &PixelDeleteInput{
		Date:    String(pixelDate),
		GraphID: String(graphID),
	}
	result, err := e2eClient.Pixel().Delete(input)
	if err != nil {
		t.Errorf("Pixel.Delete() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Pixel.Delete() got: %+v\nwant: true", result)
	}
}
