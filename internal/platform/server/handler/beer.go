package handler

import (
	"encoding/json"
	"errors"
	"falabella/internal/domain"
	"falabella/internal/platform/server/handler/model"
	"falabella/internal/usecase"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/httplog"

	"github.com/go-chi/chi/v5"
)

type Beer struct {
	findBeer     *usecase.FindBeer
	createBeer   *usecase.CreateBeer
	boxPriceBeer *usecase.BoxPriceBeer
}

func NewBeer(findBeer *usecase.FindBeer, createBeer *usecase.CreateBeer, boxPriceBeer *usecase.BoxPriceBeer) *Beer {
	return &Beer{
		findBeer:     findBeer,
		createBeer:   createBeer,
		boxPriceBeer: boxPriceBeer,
	}
}

func (b *Beer) Get(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())

	beers, err := b.findBeer.FindAll()
	if err != nil {
		log.Printf("Failed find all: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := make([]*model.Beer, len(beers))
	for i, b := range beers {
		result[i] = model.NewBeer(&b)
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		log.Err(err).Msg("failed marshall response")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (b *Beer) Post(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())

	beer := model.Beer{}

	err := json.NewDecoder(r.Body).Decode(&beer)
	if err != nil {
		log.Err(err).Msg("failed decoder body")
		http.Error(w, "Request invalid,failed decoded body", http.StatusBadRequest)
		return
	}

	id, err := b.createBeer.CreateBeer(domain.NewBeer(beer.ID, beer.Name, beer.Brewery, beer.Country, beer.Price, beer.Currency))
	if err != nil {
		if err.Error() == "beer ID already exists" {
			http.Error(w, "Beer ID already exists", http.StatusConflict)
			return
		}
		log.Err(err).Msg("failed create beer")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("Beer created id: %d", id)))
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *Beer) GetBeerxId(w http.ResponseWriter, r *http.Request) {
	beerID, err := strconv.Atoi(chi.URLParam(r, "beerID"))
	if err != nil {
		log.Printf("Failed decode beer id: %v", err)
		http.Error(w, "the beerID is required", http.StatusBadRequest)
		return
	}
	modelBeer, err := b.findBeer.FindById(beerID)
	if err != nil {
		log.Printf("Failed FindById x beer id: %v", err)
		http.Error(w, "the Failed FindById", http.StatusBadRequest)
		return
	}
	if modelBeer.ID() == 0 {
		http.Error(w, "the beer ID does not exist ", http.StatusNotFound)
		return
	}

	log := httplog.LogEntry(r.Context())

	result := model.NewBeer(modelBeer)

	marshal, err := json.Marshal(result)
	if err != nil {
		log.Err(err).Msg("Failed Json Marshall all")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (b *Beer) BoxPrice(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())

	beerID, err := strconv.Atoi(chi.URLParam(r, "beerID"))
	if err != nil {
		http.Error(w, "the currency is required", http.StatusBadRequest)
		return
	}

	cur := r.URL.Query().Get("currency")
	if cur == "" {
		http.Error(w, "the currency is required", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		quantity = 6
	}

	total, err := b.boxPriceBeer.Calculate(beerID, cur, quantity)
	if err == nil {

		_, err = w.Write([]byte(fmt.Sprintf("Operación exitosa %v", total)))
		if err != nil {
			log.Err(err).Msg("failed write body")
			return
		}

		return
	}

	if errors.Is(err, usecase.ErrBeerNotExists) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Err(err).Msg("failed calculating box price")
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
