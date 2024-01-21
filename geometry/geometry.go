package geometry

import (
	"math"
)

type FTransform struct {
	Rot         FQuat
	Translation FVector
	Pad         [4]byte
	Scale       FVector
	Pad1        [4]byte
}

type FQuat struct {
	X, Y, Z, W float32
}

type FVector struct {
	X float32
	Y float32
	Z float32
}

type D3DMATRIX struct {
	V_11, V_12, V_13, V_14 float32
	V_21, V_22, V_23, V_24 float32
	V_31, V_32, V_33, V_34 float32
	V_41, V_42, V_43, V_44 float32
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

func (t *FTransform) ToMatrixWithScale() D3DMATRIX {
	var m D3DMATRIX

	m.V_41 = t.Translation.X
	m.V_42 = t.Translation.Y
	m.V_43 = t.Translation.Z

	x2 := t.Rot.X + t.Rot.X
	y2 := t.Rot.Y + t.Rot.Y
	z2 := t.Rot.Z + t.Rot.Z

	xx2 := t.Rot.X * x2
	yy2 := t.Rot.Y * y2
	zz2 := t.Rot.Z * z2
	m.V_11 = (1.0 - (yy2 + zz2)) * t.Scale.X
	m.V_22 = (1.0 - (xx2 + zz2)) * t.Scale.Y
	m.V_33 = (1.0 - (xx2 + yy2)) * t.Scale.Z

	yz2 := t.Rot.Y * z2
	wx2 := t.Rot.W * x2
	m.V_32 = (yz2 - wx2) * t.Scale.Z
	m.V_23 = (yz2 + wx2) * t.Scale.Y

	xy2 := t.Rot.X * y2
	wz2 := t.Rot.W * z2
	m.V_21 = (xy2 - wz2) * t.Scale.Y
	m.V_12 = (xy2 + wz2) * t.Scale.X

	xz2 := t.Rot.X * z2
	wy2 := t.Rot.W * y2
	m.V_31 = (xz2 + wy2) * t.Scale.Z
	m.V_13 = (xz2 - wy2) * t.Scale.X

	m.V_14 = 0.0
	m.V_24 = 0.0
	m.V_34 = 0.0
	m.V_44 = 1.0

	return m
}

func MatrixMultiplication(pM1, pM2 D3DMATRIX) D3DMATRIX {
	var pOut D3DMATRIX

	pOut.V_11 = pM1.V_11*pM2.V_11 + pM1.V_12*pM2.V_21 + pM1.V_13*pM2.V_31 + pM1.V_14*pM2.V_41
	pOut.V_12 = pM1.V_11*pM2.V_12 + pM1.V_12*pM2.V_22 + pM1.V_13*pM2.V_32 + pM1.V_14*pM2.V_42
	pOut.V_13 = pM1.V_11*pM2.V_13 + pM1.V_12*pM2.V_23 + pM1.V_13*pM2.V_33 + pM1.V_14*pM2.V_43
	pOut.V_14 = pM1.V_11*pM2.V_14 + pM1.V_12*pM2.V_24 + pM1.V_13*pM2.V_34 + pM1.V_14*pM2.V_44
	pOut.V_21 = pM1.V_21*pM2.V_11 + pM1.V_22*pM2.V_21 + pM1.V_23*pM2.V_31 + pM1.V_24*pM2.V_41
	pOut.V_22 = pM1.V_21*pM2.V_12 + pM1.V_22*pM2.V_22 + pM1.V_23*pM2.V_32 + pM1.V_24*pM2.V_42
	pOut.V_23 = pM1.V_21*pM2.V_13 + pM1.V_22*pM2.V_23 + pM1.V_23*pM2.V_33 + pM1.V_24*pM2.V_43
	pOut.V_24 = pM1.V_21*pM2.V_14 + pM1.V_22*pM2.V_24 + pM1.V_23*pM2.V_34 + pM1.V_24*pM2.V_44
	pOut.V_31 = pM1.V_31*pM2.V_11 + pM1.V_32*pM2.V_21 + pM1.V_33*pM2.V_31 + pM1.V_34*pM2.V_41
	pOut.V_32 = pM1.V_31*pM2.V_12 + pM1.V_32*pM2.V_22 + pM1.V_33*pM2.V_32 + pM1.V_34*pM2.V_42
	pOut.V_33 = pM1.V_31*pM2.V_13 + pM1.V_32*pM2.V_23 + pM1.V_33*pM2.V_33 + pM1.V_34*pM2.V_43
	pOut.V_34 = pM1.V_31*pM2.V_14 + pM1.V_32*pM2.V_24 + pM1.V_33*pM2.V_34 + pM1.V_34*pM2.V_44
	pOut.V_41 = pM1.V_41*pM2.V_11 + pM1.V_42*pM2.V_21 + pM1.V_43*pM2.V_31 + pM1.V_44*pM2.V_41
	pOut.V_42 = pM1.V_41*pM2.V_12 + pM1.V_42*pM2.V_22 + pM1.V_43*pM2.V_32 + pM1.V_44*pM2.V_42
	pOut.V_43 = pM1.V_41*pM2.V_13 + pM1.V_42*pM2.V_23 + pM1.V_43*pM2.V_33 + pM1.V_44*pM2.V_43
	pOut.V_44 = pM1.V_41*pM2.V_14 + pM1.V_42*pM2.V_24 + pM1.V_43*pM2.V_34 + pM1.V_44*pM2.V_44

	return pOut
}

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

	matrix := D3DMATRIX{}
	matrix.V_11 = CP * CY
	matrix.V_12 = SR*SP*CY - CR*SY
	matrix.V_13 = -(CR*SP*CY + SR*SY)
	matrix.V_14 = origin.X

	matrix.V_21 = CP * SY
	matrix.V_22 = SR*SP*SY + CR*CY
	matrix.V_23 = CY*SR - CR*SP*SY
	matrix.V_24 = origin.Y

	matrix.V_31 = SP
	matrix.V_32 = -SR * CP
	matrix.V_33 = CR * CP
	matrix.V_34 = origin.Z

	matrix.V_41 = 0.0
	matrix.V_42 = 0.0
	matrix.V_43 = 0.0
	matrix.V_44 = 1.0

	return matrix
}

func ProjectWorldToScreen(worldLocation FVector, viewInfo FMinimalViewInfo) FVector {
	screenLocation := FVector{0, 0, 0}

	cameraLocation := viewInfo.Location
	cameraRotation := viewInfo.Rotation

	tempMatrix := Matrix(cameraRotation, FVector{0, 0, 0})

	vAxisX := FVector{tempMatrix.V_11, tempMatrix.V_21, tempMatrix.V_31}
	vAxisY := FVector{tempMatrix.V_12, tempMatrix.V_22, tempMatrix.V_32}
	vAxisZ := FVector{tempMatrix.V_13, tempMatrix.V_23, tempMatrix.V_33}

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
