package geometry

import "math"

type FVector struct {
	X float32
	Y float32
	Z float32
}

type FMinimalViewInfo struct {
	Location FVector
	Rotation FVector
	FOV      float32
}

func NewFVector(x, y, z float32) FVector {
	return FVector{X: x, Y: y, Z: z}
}

func (v FVector) Dot(other FVector) float32 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v FVector) Distance(other FVector) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return float32(math.Sqrt(float64(dx*dx+dy*dy+dz*dz))) * 0.03048
}

func (v FVector) Add(other FVector) FVector {
	return FVector{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v FVector) Subtract(other FVector) FVector {
	return FVector{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v FVector) Multiply(number float32) FVector {
	return FVector{X: v.X * number, Y: v.Y * number, Z: v.Z * number}
}

func (v FVector) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v FVector) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v FVector) Normalize() FVector {
	vector := FVector{}
	length := v.Magnitude()

	if length != 0 {
		vector.X = v.X / length
		vector.Y = v.Y / length
		vector.Z = v.Z / length
	} else {
		vector.X, vector.Y = 0.0, 0.0
		vector.Z = 1.0
	}

	return vector
}

func (v *FVector) AddAssign(other FVector) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}
