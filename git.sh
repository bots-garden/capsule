#!/bin/bash

message=""
case $1 in

  # 🎨: art
  art)
    message="Improve structure / format of the code"
    emoji="🎨"
    ;;

  # 🐛: bug
  bug|fix)
    message="Fix a bug"
    emoji="🐛"
    ;;

  # ✨: sparkles
  sparkles)
    message="Introduce new features"
    emoji="✨"
    ;;

  # 📝: memo
  memo|doc|documentation)
    message="Add or update documentation"
    emoji="📝"
    ;;

  # 🌸: cherry_blossom
  gardening|garden|clean|cleaning)
    message="Gardening"
    emoji="🌸"
    ;;


  # 🚀: rocket
  rocket|deploy)
    message="Deploy stuff"
    emoji="🚀"
    ;;


  # 🎉: tada
  tada|first)
    message="Begin a project"
    emoji="🎉"
    ;;

  # 🚧: construction
  construction|wip)
    message="Work in progress"
    emoji="🚧"
    ;;

  # 📦️: package
  package|build)
    message="Add or update compiled files or packages"
    emoji="📦️"
    ;;

  # 👽️: alien
  alien|api)
    message="Update code due to external API changes"
    emoji="👽️"
    ;;

  # 🐳: whale
  docker|container)
    message="Docker"
    emoji="🐳"
    ;;

  # ⚗️: alembic
  alembic|experiments|experiment|xp)
    message="Perform experiments"
    emoji="⚗"
    ;;


  # 💾: floppy-disk
  save)
    message="Saved"
    emoji="💾"
    ;;

  *)
  message="Updated"
  emoji="🛟"
  ;;

esac

if [ -z "$2" ]
then
  # empty
  git add .; git commit -m "$emoji $message."; git push
else
  # not empty
  git add .; git commit -m "$emoji $message: $2"; git push
fi
