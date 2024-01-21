package geometry

import (
	"math"
)

type FVector struct {
	X float32
	Y float32
	Z float32
}

type D3DMATRIX struct {
	M [4][4]float32
}

type FMinimalViewInfo struct {
	Location FVector
	Rotation FVector
	FOV      float32
}

const (
	Width  = 1920
	Height = 1080
)

func Matrix(rot FVector, origin FVector) D3DMATRIX {
	radPitch := (rot.X * float32(math.Pi) / 180.0)
	radYaw := (rot.Y * float32(math.Pi) / 180.0)
	radRoll := (rot.Z * float32(math.Pi) / 180.0)

	SP := float32(math.Sin(float64(radPitch)))
	CP := float32(math.Cos(float64(radPitch)))
	SY := float32(math.Sin(float64(radYaw)))
	CY := float32(math.Cos(float64(radYaw)))
	SR := float32(math.Sin(float64(radRoll)))
	CR := float32(math.Cos(float64(radRoll)))

	matrix := D3DMATRIX{
		M: [4][4]float32{
			{CP * CY, CP * SY, SP, 0.0},
			{SR*SP*CY - CR*SY, SR*SP*SY + CR*CY, -SR * CP, 0.0},
			{-(CR*SP*CY + SR*SY), CY*SR - CR*SP*SY, CR * CP, 0.0},
			{origin.X, origin.Y, origin.Z, 1.0},
		},
	}

	return matrix
}

func ProjectWorldToScreen(worldLocation FVector, viewInfo FMinimalViewInfo) FVector {
	screenLocation := FVector{0, 0, 0}

	cameraLocation := viewInfo.Location
	cameraRotation := viewInfo.Rotation

	tempMatrix := Matrix(cameraRotation, FVector{0, 0, 0})

	vAxisX := FVector{tempMatrix.M[0][0], tempMatrix.M[0][1], tempMatrix.M[0][2]}
	vAxisY := FVector{tempMatrix.M[1][0], tempMatrix.M[1][1], tempMatrix.M[1][2]}
	vAxisZ := FVector{tempMatrix.M[2][0], tempMatrix.M[2][1], tempMatrix.M[2][2]}

	vDelta := worldLocation.Subtract(cameraLocation)
	vTransformed := FVector{
		X: vDelta.Dot(vAxisY),
		Y: vDelta.Dot(vAxisZ),
		Z: vDelta.Dot(vAxisX),
	}

	if vTransformed.Z < 1.0 {
		vTransformed.Z = 1.0
	}

	fovAngle := viewInfo.FOV
	screenCenterX := Width / 2.0
	screenCenterY := Height / 2.0

	screenLocation.X = float32(screenCenterX + float64(vTransformed.X)*(screenCenterX/math.Tan(float64(fovAngle)*(math.Pi/360.0)))/float64(vTransformed.Z))
	screenLocation.Y = float32(screenCenterY - float64(vTransformed.Y)*(screenCenterX/math.Tan(float64(fovAngle)*(math.Pi/360.0)))/float64(vTransformed.Z))
	return screenLocation
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
