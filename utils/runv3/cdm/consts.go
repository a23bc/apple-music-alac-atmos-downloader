package wv

import "encoding/base64"

var DefaultPrivateKey string
var DefaultClientID []byte

func InitConstants() {

	DefaultPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA2bO3yvFwNnIHsbDl3MTjKdDsiBWsuZWOGVxInFWAVMp+nffG\nYlquTKpJurEry95yprcRB3hYhvA5ghsACidcWPDEPVqqRZ7YXLevyUA+Sn2Jxpvt\nOcwyFHbSwruNxprWOkHCT774O4L/wJUt5x2C4iFCrJByjw0omN8u+EHdavvH7ZPn\nb3/EZp/cpZa9/+HOkutvBHBvaPp18F8JQhzUQ9MwLuDFTr+QLDB5+Y57Je2tNYDK\nxD1K+Ed5Ja0A4OKhPKIwPwPre0nt5scjLba3LSAKtKxiGqFtWO4U7Tf1YrdjJv2o\n9o8Sf8qcnbpzvQ4KwFqehuJnB7+W7mdJJw12PQIDAQABAoIBACE32wOMc6LbI3Fp\nnKljIYZv6qeZJxHqUBRukGXKZhqKC2fvNsYrMA1irn1eK2CgQL5PkLmjE18DqMLB\ne/AQsXagxlDWVMTqx/jdzmTW+KpFHZDAmiIHllypBN/R3oA/gBDDl/KzIQ1zn7Kz\nEJ4DUsVObe4G3HQXfepVo8Udx7tbB7X6wHe2kEgFyY3lPdvubik0C4t4ipSD79y7\nSfW7XVA5XUQmqN4U2kWM0uSwzd4BA7hqyScJsygf6KgpMWPS2xFZEZQRUpYcBH48\nE7YqNrrlYP3yaQ+9Jx56kKS0mvv3vUXS7AfUbU8CiHwD9I3BGwswEUueOGGVeXbx\ntFF8s8ECgYEA97BDcL/bt+r3qJF0dxtMB5ZngJbFx9RdsblYepVpblr2UfxnFttO\nPoNSKa4W36HuDsun49dkaoABJWdtZs2Hy6q+xvEgozvhMaBVE3spnWnzCT1yTMYL\nG02uDEl0dPiTg116bVElaswtqMXvnnpbOTMTe7Ig9sWiUW/GH9RM+N8CgYEA4QHb\n+OA0BfczbVQP9B+plt4mAuu4BDm4GPwq1yXOWo3Ct8Ik+HeY1hqOObpfyQMAza+E\ne/kP6W8vXpiElGrmiUbTXK4Rzmf+yYeOrvl3D80bFq4GtDNAIQD3jpj6zjlT+Gzw\nI501gRx5iPl4fSccRSdpoeri7F9ANtc6EEGFyGMCgYEAjMznWYXHGkL47BtbkIW0\n769BQSj0X4dKh8gsEusylugglDSeSbD7RrASGd175T7A/CorU2rTC3OesyubVlBJ\n/K4gaykRe5mDh1l0Y3GlE3XyEXObsSb3k1rSMOvkxsWz3X5bJR923MIaxpFWiMlX\naCmvzqZQ9NceUZrvjpJ5+xMCgYAJa8KCESEcftUwZqykVA8Nug9tX+E8jA4hPa2t\nhG+3augUOZTCsn87t7Dsydjo2a9W7Vpmtm7sHzOkik5CyJcOeGCxKLimI8SPO5XF\nzbwmdTgFIxQ0x1CQETJMTityJwRVCnqjgxmSZlbQXWGmG9UbMCNEHEmUDAjsQuaz\nd4racQKBgQDR1Y2kalvleYGrhwcA8LTnIh0rYEfAt9YxNmTi5qDKf5QPvUP2v+WO\nfSB5coUqR8LBweHE5V8JgFt74fdLBqZV/k2z/dI0r+EQWmpZ2uPEC0Khk/Sb9iRD\nfH7at3PMusrkwZCGZ8beFEAr6icXclV08nPCNGB6WckacfzpAj8Azg==\n-----END RSA PRIVATE KEY-----"
	DefaultClientdIDBase64 := "CAESmgsK3QMIAhIQeeRrycR5oAnVvSCrdzFrTxivgsKlBiKOAjCCAQoCggEBANmzt8rxcDZyB7Gw5dzE4ynQ7IgVrLmVjhlcSJxVgFTKfp33xmJarkyqSbqxK8vecqa3EQd4WIbwOYIbAAonXFjwxD1aqkWe2Fy3r8lAPkp9icab7TnMMhR20sK7jcaa1jpBwk+++DuC/8CVLecdguIhQqyQco8NKJjfLvhB3Wr7x+2T529/xGaf3KWWvf/hzpLrbwRwb2j6dfBfCUIc1EPTMC7gxU6/kCwwefmOeyXtrTWAysQ9SvhHeSWtAODioTyiMD8D63tJ7ebHIy22ty0gCrSsYhqhbVjuFO039WK3Yyb9qPaPEn/KnJ26c70OCsBanobiZwe/lu5nSScNdj0CAwEAASjwIkgBUqoBCAEQABqBAQQZhh0LPs5wmuuobaJofVK1k0DjvnNhqvOMfGw0Zlzum4aTAvasMiyWfhjo/+xmHtsRvK3ek9EOdIB1e2c5azFuScAMS2n7ZGzqA8XBb+UPM46FUeGt7o1jDm/AysaZt4U6Ji8wXl41dWA9kF/iIK7uThSmb+mhspLLYo3AUiu2hiIgFm8idU4+UvSfVB4JveJ+hqeNbpYuNWkrxlbj9DDjWgYSgAIemDQcy+RKUwwGq59NhaxYSH3hxSHGCkhcXnjNC0OeV5gBdJQl7uqN90lkF3JxnlvYF3mhux7pZR5jii4KaNG6+vZXEq21irNMnoSxwIlzvpMov7xOvQWVm00K+xDkO20ncTC1ClXpmAAHyDXmMeTrzvCLo7tc3USbaImlIWAX92saZojzJ3n9gc+cjBKGqz2AgcsFCigSZ5vpLtz/wEk5PxIGKJ6OWjEy4D5HZG0p2MYyhM84fUh3TOfuexK1ceWrOfPxCbxSPRi9w0BEaDmixt/K4mIalUFTBJsWxtE6ww38UmFLktWoMM8+QLnhxe6jmuVpuchdLtnMPnkAs6XjGrQFCq4CCAESEGnj6Ji7LD+4o7MoHYT4jBQYjtW+kQUijgIwggEKAoIBAQDY9um1ifBRIOmkPtDZTqH+CZUBbb0eK0Cn3NHFf8MFUDzPEz+emK/OTub/hNxCJCao//pP5L8tRNUPFDrrvCBMo7Rn+iUb+mA/2yXiJ6ivqcN9Cu9i5qOU1ygon9SWZRsujFFB8nxVreY5Lzeq0283zn1Cg1stcX4tOHT7utPzFG/ReDFQt0O/GLlzVwB0d1sn3SKMO4XLjhZdncrtF9jljpg7xjMIlnWJUqxDo7TQkTytJmUl0kcM7bndBLerAdJFGaXc6oSY4eNy/IGDluLCQR3KZEQsy/mLeV1ggQ44MFr7XOM+rd+4/314q/deQbjHqjWFuVr8iIaKbq+R63ShAgMBAAEo8CISgAMii2Mw6z+Qs1bvvxGStie9tpcgoO2uAt5Zvv0CDXvrFlwnSbo+qR71Ru2IlZWVSbN5XYSIDwcwBzHjY8rNr3fgsXtSJty425djNQtF5+J2jrAhf3Q2m7EI5aohZGpD2E0cr+dVj9o8x0uJR2NWR8FVoVQSXZpad3M/4QzBLNto/tz+UKyZwa7Sc/eTQc2+ZcDS3ZEO3lGRsH864Kf/cEGvJRBBqcpJXKfG+ItqEW1AAPptjuggzmZEzRq5xTGf6or+bXrKjCpBS9G1SOyvCNF1k5z6lG8KsXhgQxL6ADHMoulxvUIihyPY5MpimdXfUdEQ5HA2EqNiNVNIO4qP007jW51yAeThOry4J22xs8RdkIClOGAauLIl0lLA4flMzW+VfQl5xYxP0E5tuhn0h+844DslU8ZF7U1dU2QprIApffXD9wgAACk26Rggy8e96z8i86/+YYyZQkc9hIdCAERrgEYCEbByzONrdRDs1MrS/ch1moV5pJv63BIKvQHGvLkaFgoMY29tcGFueV9uYW1lEgZHb29nbGUaIQoKbW9kZWxfbmFtZRITQU9TUCBvbiBJQSBFbXVsYXRvchoYChFhcmNoaXRlY3R1cmVfbmFtZRIDeDg2Gh4KC2RldmljZV9uYW1lEg9nZW5lcmljX3g4Nl9hcm0aIgoMcHJvZHVjdF9uYW1lEhJzZGtfZ3Bob25lX3g4Nl9hcm0aZAoKYnVpbGRfaW5mbxJWZ29vZ2xlL3Nka19ncGhvbmVfeDg2X2FybS9nZW5lcmljX3g4Nl9hcm06OS9QU1IxLjE4MDcyMC4xMjIvNjczNjc0Mjp1c2VyZGVidWcvZGV2LWtleXMaHgoUd2lkZXZpbmVfY2RtX3ZlcnNpb24SBjE0LjAuMBokCh9vZW1fY3J5cHRvX3NlY3VyaXR5X3BhdGNoX2xldmVsEgEwMg4QASAAKA0wAEAASABQAA=="
	DefaultClientID, _ = base64.StdEncoding.DecodeString(DefaultClientdIDBase64)

	// DefaultPrivateKeyBuffer, err := ioutil.ReadFile("device_private_key")
	// if err != nil {
	// 	panic(err)
	// }
	// DefaultPrivateKey = string(DefaultPrivateKeyBuffer)
 //
	// DefaultClientID, err = ioutil.ReadFile("device_client_id_blob")
	// if err != nil {
	// 	panic(err)
	// }
}
