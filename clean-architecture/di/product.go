package di

import (
	productService "clean-architecture/adaptar/http/service/product"
	"clean-architecture/adaptar/postgres"
	productRepository "clean-architecture/adaptar/postgres/repository/product"
	"clean-architecture/core/domain"
	productUseCase "clean-architecture/core/domain/usecase/product"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productRepository.New(conn)
	productUseCase := productUseCase.New(productRepository)
	productServide := productService.New(productUseCase)

	return productServide
}
