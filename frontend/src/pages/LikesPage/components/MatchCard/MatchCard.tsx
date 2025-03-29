import React from "react";
import styles from "./matchCard.module.css";
// import { IconButton } from "@mui/material";
import { TelegramIcon } from "../../../../shared/components/TelegramIcon";
import { InstagramIcon } from "../../../../shared/components/InstagramIcon";
import { IconButton } from "../../../../shared/components/IconButton";

// Define the match profile interface
export interface MatchProfile {
  id: number;
  name: string;
  age: number;
  imageUrl: string;
  telegram: string;
  instagram: string;
}

interface MatchCardProps {
  match: MatchProfile;
  onClick: (match: MatchProfile) => void;
  onSocialClick: (matchId: number, type: "telegram" | "instagram") => void;
  showChatMessage: number | null;
}

export const MatchCard: React.FC<MatchCardProps> = ({
  match,
  onClick,
  onSocialClick,
  showChatMessage,
}) => {
  return (
    <div className={styles.matchCard} onClick={() => onClick(match)}>
      <img
        src={match.imageUrl}
        alt={match.name}
        className={styles.matchImage}
      />
      <div className={styles.matchInfo}>
        <div className={styles.matchName}>
          {match.name}, {match.age}
        </div>
        {showChatMessage === match.id && (
          <div className={styles.chatMessage}>Открыто в новой вкладке</div>
        )}
      </div>
      <div
        className={styles.socialButtons}
        onClick={(e) => e.stopPropagation()}
      >
        <IconButton
          // size="small"
          className={styles.socialButton}
          onClick={() => onSocialClick(match.id, "telegram")}
        >
          <TelegramIcon />
        </IconButton>
        <IconButton
          // size="small"
          className={styles.socialButton}
          onClick={() => onSocialClick(match.id, "instagram")}
        >
          <InstagramIcon />
        </IconButton>
      </div>
    </div>
  );
};

export default MatchCard;
